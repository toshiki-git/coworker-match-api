package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IMessageRepo interface {
	GetMessages(userId, matchingId string) (*models.GetMessageRes, error)
	GetLastQuestionCardText(matchingId string) (*string, error)
	CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error)
	UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error)
}

type messageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) IMessageRepo {
	return &messageRepo{db: db}
}

func (mr *messageRepo) GetMessages(userId, matchingId string) (*models.GetMessageRes, error) {
	query := `
		SELECT 
			qc.question_card_id,
			qc.question_card_text,
			me_msg.message_id AS me_message_id,
			me_msg.message_text AS me_message_text,
			you_msg.message_id AS you_message_id,
			you_msg.message_text AS you_message_text
		FROM 
			matchings m
		JOIN (
			SELECT
				message_text,
				matching_id,
				question_card_id,
				message_id
			FROM
				messages
			WHERE
				user_id = $1
			) me_msg ON m.matching_id = me_msg.matching_id
		JOIN (
			SELECT
				message_text,
				matching_id,
				question_card_id,
				message_id
			FROM
				messages
			WHERE
				user_id != $1
			) you_msg ON m.matching_id = you_msg.matching_id
		JOIN question_cards qc ON me_msg.question_card_id = qc.question_card_id AND you_msg.question_card_id = qc.question_card_id
		WHERE 
			m.matching_id = $2;
		`
	response := models.NewGetMessageRes([]models.GetMessageResMessagesInner{})
	rows, err := mr.db.Query(query, userId, matchingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data models.GetMessageResMessagesInner
		me := *models.NewMessage("", "")
		you := *models.NewMessage("", "")
		data.MessagePair = *models.NewGetMessageResMessagesInnerMessagePair(me, you)

		if err := rows.Scan(&data.QuestionCardId, &data.QuestionCardText, &data.MessagePair.Me.MessageId, &data.MessagePair.Me.MessageText, &data.MessagePair.You.MessageId, &data.MessagePair.You.MessageText); err != nil {
			return nil, err
		}
		response.Messages = append(response.Messages, data)
	}

	// メッセージをis_readに更新するクエリ
	updateQuery := `
			UPDATE messages 
			SET is_read = TRUE 
			WHERE matching_id = $1 
			AND user_id != $2 
			AND is_read = FALSE
		`

	if _, err := mr.db.Exec(updateQuery, matchingId, userId); err != nil {
		return nil, err
	}

	return response, nil
}

func (mr *messageRepo) GetLastQuestionCardText(matchingId string) (*string, error) {
	query := `
			SELECT
				COALESCE(qc.question_card_text, 'メッセージはありません')
			FROM
				messages msg
			JOIN
				question_cards qc ON msg.question_card_id = qc.question_card_id
			WHERE
				msg.matching_id = $1
			ORDER BY
				msg.created_at DESC
			LIMIT 1;
			`
	var questionCardText string
	if err := mr.db.QueryRow(query, matchingId).Scan(&questionCardText); err != nil {
		return nil, err
	}

	return &questionCardText, nil
}

func (mr *messageRepo) CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error) {
	var otherUserId string
	query := `
		SELECT 
			CASE 
				WHEN sender_user_id = $1 THEN receiver_user_id
				ELSE sender_user_id
			END AS other_user_id
		FROM 
			matchings
		WHERE 
			matching_id = $2`

	if err := mr.db.QueryRow(query, userId, matchingId).Scan(&otherUserId); err != nil {
		return nil, err
	}

	tx, err := mr.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var response models.CreateMessageRes

	// 自分のIDでメッセージを作成
	err = tx.QueryRow(`
		INSERT INTO messages (matching_id, question_card_id, user_id)
		VALUES ($1, $2, $3)
		RETURNING message_id`, matchingId, req.QuestionCardId, userId).Scan(&response.MessageId)
	if err != nil {
		return nil, err
	}

	// 相手のIDでメッセージを作成
	_, err = tx.Exec(`
		INSERT INTO messages (matching_id, question_card_id, user_id)
		VALUES ($1, $2, $3)`, matchingId, req.QuestionCardId, otherUserId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (mr *messageRepo) UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error) {
	query := `
		UPDATE messages 
		SET message_text = $1, updated_at = NOW() 
		WHERE message_id = $2
	`

	_, err := mr.db.Exec(query, req.MessageText, messageId)
	if err != nil {
		return nil, err
	}

	response := &models.UpdateMessageRes{
		MessageText: req.MessageText,
	}

	return response, nil
}
