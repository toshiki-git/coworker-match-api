package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IMessageRepo interface {
	GetMessages(userId, matchingId string) (*models.GetMessageRes, error)
	UpdateMessagesReadStatus(matchingId, userId string) error
	CreateMessage(userId, otherUserId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error)
	GetOtherUserId(userId, matchingId string) (string, error)
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
			me.message_id AS me_message_id,
			me.message_text AS me_message_text,
			you.message_id AS you_message_id,
			you.message_text AS you_message_text
		FROM 
			matchings m
		JOIN messages me ON m.matching_id = me.matching_id AND me.user_id = $1
		JOIN messages you ON m.matching_id = you.matching_id AND you.user_id != $1
		JOIN question_cards qc ON me.question_card_id = qc.question_card_id AND you.question_card_id = qc.question_card_id
		WHERE 
			m.matching_id = $2
		ORDER BY
			me.created_at ASC;
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

	return response, nil
}

func (mr *messageRepo) UpdateMessagesReadStatus(matchingId, userId string) error {
	query := `
			UPDATE 
				messages 
			SET 
				is_read = TRUE 
			WHERE
				matching_id = $1 
				AND user_id != $2 
				AND is_read = FALSE
			`

	if _, err := mr.db.Exec(query, matchingId, userId); err != nil {
		return err
	}

	return nil
}

func (mr *messageRepo) CreateMessage(userId, otherUserId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error) {
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
	response.MessageId, err = mr.InsertMessage(tx, matchingId, req.QuestionCardId, userId)
	if err != nil {
		return nil, err
	}

	// 相手のIDでメッセージを作成
	_, err = mr.InsertMessage(tx, matchingId, req.QuestionCardId, otherUserId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (mr *messageRepo) GetOtherUserId(userId, matchingId string) (string, error) {
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
		return "", err
	}

	return otherUserId, nil
}

func (mr *messageRepo) InsertMessage(tx *sql.Tx, matchingId, questionCardId, userId string) (string, error) {
	query := `
			INSERT INTO messages (
				matching_id,
				question_card_id,
				user_id
			)
			VALUES (
				$1,
				$2,
				$3
			)
			RETURNING
				message_id
			`

	var messageId string
	err := tx.QueryRow(query, matchingId, questionCardId, userId).Scan(&messageId)
	if err != nil {
		return "", err
	}
	return messageId, nil
}

func (mr *messageRepo) UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error) {
	query := `
		UPDATE
			messages 
		SET 
			message_text = $1,
			updated_at = NOW()
		WHERE 
			message_id = $2
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
