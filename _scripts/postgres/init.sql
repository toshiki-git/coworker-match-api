-- =============================================
-- ========== テーブルの定義 ==========
-- =============================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(255) PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    avatar_url VARCHAR(255) NOT NULL,
    age SMALLINT DEFAULT 0,
    gender VARCHAR(6) DEFAULT 'other' CHECK (gender IN ('male', 'female', 'other')),
    birthplace VARCHAR(255) DEFAULT '',
    job_type VARCHAR(255) DEFAULT '',
    line_account VARCHAR(255) DEFAULT '',
    discord_account VARCHAR(255) DEFAULT '',
    biography TEXT DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    category_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS hobbies (
    hobby_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    hobby_name VARCHAR(255) NOT NULL,
    category_id UUID NOT NULL,
    creator_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (creator_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_hobbies (
    hobby_id UUID NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (hobby_id, user_id),
    FOREIGN KEY (hobby_id) REFERENCES hobbies(hobby_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS matching_questions (
    question_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    question_text TEXT NOT NULL,
    question_category_id UUID,
    FOREIGN KEY (question_category_id) REFERENCES categories(category_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS matches (
    matching_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sender_user_id VARCHAR(255) NOT NULL,
    receiver_user_id VARCHAR(255) NOT NULL,
    status VARCHAR(8) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'rejected')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (sender_user_id, receiver_user_id),
    FOREIGN KEY (sender_user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS question_cards (
    question_card_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    question_card_text TEXT NOT NULL,
    creator_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (creator_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages (
    message_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    matching_id UUID NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    question_card_id UUID NOT NULL,
    message_text TEXT DEFAULT NULL,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (matching_id, question_card_id, user_id),
    FOREIGN KEY (matching_id) REFERENCES matches(matching_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (question_card_id) REFERENCES question_cards(question_card_id) ON DELETE CASCADE
);


-- =============================================
-- ========== これより下はデフォルトデータ ==========
-- =============================================

-- -------------------------------------------
-- users のデフォルトデータ
-- -------------------------------------------
INSERT INTO users (user_id, user_name, email, avatar_url, age, gender, birthplace, job_type, line_account, discord_account, biography)
VALUES
('taro_sato_id', '佐藤 太郎', 'taro.sato@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhbP_zxvaiaDB8eqIzrvyz-U-_UUT7aoasiBkhvJcMbh7qz7Dr0oHiCmVzpyYMBLigs7SwlKRhcvYLRr37rcX5eOzTSC8zcHvhM_e3j8TxQDFkmVt93h6eh1sxckKV4jLQtZO1_KevnhxQ/s800/animal_mark01_buta.png', 25, 'male', '東京', 'エンジニア', 'taro_sato', 'taro_sato_discord', '趣味はプログラミングです。'),
('jiro_suzuki_id', '鈴木 次郎', 'jiro.suzuki@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh_ri_MUkJIyJyysAHkokKUHVMj3WNV8BolUAavT8dB_xo61dWlvF4K6oO2x9wdtPVKGWCpYrXQgKDtU_u2XPBhu8OOmwod0OpLonPKCxhwZbaLnKxh_VgC6H3kbeyU7_DqXP-ZbEkIcVo/s800/animal_mark02_kuma.png', 30, 'male', '大阪', 'デザイナー', 'jiro_suzuki', 'jiro_suzuki_discord', 'デザインが大好きです。'),
('hanako_takahashi_id', '高橋 花子', 'hanako.takahashi@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgAbQzjvVfLxjuDrwBeXfXwrmWZlGXxETik8r1wV55AT-LFusq06xmTl5Dwraxc09Ur_7mjuojd2DuU6_KcLss36W4h26_TeaN7qmYbapafNyDpa6xyXgPJH9xCtFTbRf8DN8n_u_yaRk8/s800/animal_mark03_inu.png', 28, 'female', '北海道', 'マーケティング', 'hanako_takahashi', 'hanako_takahashi_discord', 'マーケティングの経験が豊富です。'),
('shiro_tanaka_id', '田中 四郎', 'shiro.tanaka@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEglI1wsFlukEv0MdN6TtbwaHfplbo2Hii32hm73P4hrSKdGOsirgCdqk2lHSv52eCIShi5_omhCPiPTd_Vd2yo7848-7Fb3FkA5IGAe0St8kz0yY0x8PKq2W5W9zxm-xjB0Un2TUwlO16E/s800/animal_mark04_neko.png', 35, 'male', '福岡', 'プロジェクトマネージャー', 'shiro_tanaka', 'shiro_tanaka_discord', 'プロジェクト管理が得意です。'),
('goro_ito_id', '伊藤 五郎', 'goro.ito@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiXMv2z-jXG4uJWhmrH_iMa6Pnd_sQ3FC3zqK6g3UchsmU4bUDEq8hjJOG6lwDbHJe7deikLvxM18apY9VfIpuskFxks5b72Qor5Oh-0PVf-xYlfUosNLnnZ-NcasU5EUpSprXP93DcHR8/s800/animal_mark05_zou.png', 22, 'male', '京都', 'データサイエンティスト', 'goro_ito', 'goro_ito_discord', 'データ分析が好きです。'),
('rikka_watanabe_id', '渡辺 六花', 'rikka.watanabe@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj768RPZ5b8mgZzeYuuJLVKMu3MKa7s4idSXFTBP1MPLmwvqUP0uUHzgAdrOadZiYRY4yCC8dtec0zabksNcPspqPk9mSiEHF_9Tp7Avm_UUIdL51HBm6m_GpBswyntgUpAlSRqU6gptUo/s800/animal_mark06_uma.png', 27, 'female', '広島', 'リサーチャー', 'rikka_watanabe', 'rikka_watanabe_discord', 'リサーチが得意です。'),
('nanami_yamamoto_id', '山本 七海', 'nanami.yamamoto@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgm0bF7TAkvPwpcUixv_pJH2Lm3eLRRFCUwEXTorYuAXm02VYTSbtupnXio2MOlnufSkmfaWpAFpgD0eCPGKXaVznjbZ6idvy3WPwVzGqQVLZOOGcV0d9_UHjegC9dGhMf9WSYldgXxRWA/s800/animal_mark07_lion.png', 26, 'female', '神奈川', 'ソフトウェアエンジニア', 'nanami_yamamoto', 'nanami_yamamoto_discord', 'ソフトウェア開発が得意です。'),
('hachiro_nakamura_id', '中村 八郎', 'hachiro.nakamura@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEivWfMkHHnoD3L7fHTMwqqVYQV3WYIZ5UEKtVDz8w6rgP3u_-3zN3Tatefi-sw5Ss7wuN7JU2wE8_enz_WweKpP5vR5w-WwzbLFhRxvC5hJcTcK7r1wLK9lMwhno_rdgk8RYtfPgQIsrDE/s800/animal_mark08_kaba.png', 31, 'male', '兵庫', 'プロダクトマネージャー', 'hachiro_nakamura', 'hachiro_nakamura_discord', 'プロダクト管理が得意です。'),
('kyumi_kobayashi_id', '小林 九美', 'kyumi.kobayashi@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjUm_KAhSJnmbw9teXImS5_7NqHmfWniN1KTx1VOVsPgcoknTI-FfkdIFDYySGep3Z_yIIk0RIeVT9StMhxhTPaioaW__NR5rNI7arG_lsSFYAutxSGn3J4Ib1aaFBimgSNet_jvOohBZY/s800/animal_mark10_usagi.png', 29, 'female', '千葉', '人事', 'kyumi_kobayashi', 'kyumi_kobayashi_discord', '人事管理が得意です。'),
('juro_saito_id', '斉藤 十郎', 'juro.saito@example.com', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj-ZTFMMaDsvAn_OxemB7qnLSlmba4fzlBz9wFL4NJGuDC5Q0IsMKOa8UKcBtoJk_Xu-0G3kFLBDBpaEWgvkp5up1BhkghcuEJieqd2eoRL5FDcCqpn972fBZbRu-SC1afrNDqMPS1dniQ/s800/animal_mark11_panda.png', 24, 'male', '沖縄', '営業', 'juro_saito', 'juro_saito_discord', '営業が得意です。');

-- -------------------------------------------
-- categories のデフォルトデータ
-- -------------------------------------------
INSERT INTO categories (category_id, category_name)
SELECT uuid_generate_v4(), 'インドア' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE category_name = 'インドア');
INSERT INTO categories (category_id, category_name)
SELECT uuid_generate_v4(), 'ゲーム' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE category_name = 'ゲーム');
INSERT INTO categories (category_id, category_name)
SELECT uuid_generate_v4(), '技術' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE category_name = '技術');
INSERT INTO categories (category_id, category_name)
SELECT uuid_generate_v4(), 'スポーツ' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE category_name = 'スポーツ');
INSERT INTO categories (category_id, category_name)
SELECT uuid_generate_v4(), 'アウトドア' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE category_name = 'アウトドア');
INSERT INTO categories (category_id, category_name)
SELECT uuid_generate_v4(), '音楽' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE category_name = '音楽');

-- ------------------------------------------
-- hobbies のデフォルトデータ
-- ------------------------------------------
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'アニメ', (SELECT category_id FROM categories WHERE category_name = 'インドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'アニメ');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '映画', (SELECT category_id FROM categories WHERE category_name = 'インドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '映画');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '漫画', (SELECT category_id FROM categories WHERE category_name = 'インドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '漫画');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '麻雀', (SELECT category_id FROM categories WHERE category_name = 'ゲーム')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '麻雀');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'ホラー', (SELECT category_id FROM categories WHERE category_name = 'ゲーム')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'ホラー');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'ボードゲーム', (SELECT category_id FROM categories WHERE category_name = 'ゲーム')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'ボードゲーム');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'RPG', (SELECT category_id FROM categories WHERE category_name = 'ゲーム')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'RPG');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'FPS', (SELECT category_id FROM categories WHERE category_name = 'ゲーム')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'FPS');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'パズル', (SELECT category_id FROM categories WHERE category_name = 'ゲーム')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'パズル');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'スマホゲーム', (SELECT category_id FROM categories WHERE category_name = 'ゲーム')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'スマホゲーム');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '競技プログラミング', (SELECT category_id FROM categories WHERE category_name = '技術')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '競技プログラミング');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'デザイン', (SELECT category_id FROM categories WHERE category_name = '技術')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'デザイン');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'Android開発', (SELECT category_id FROM categories WHERE category_name = '技術')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'Android開発');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'バックエンド開発', (SELECT category_id FROM categories WHERE category_name = '技術')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'バックエンド開発');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'フロントエンド開発', (SELECT category_id FROM categories WHERE category_name = '技術')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'フロントエンド開発');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'iOS開発', (SELECT category_id FROM categories WHERE category_name = '技術')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'iOS開発');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'サッカー', (SELECT category_id FROM categories WHERE category_name = 'スポーツ')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'サッカー');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '野球', (SELECT category_id FROM categories WHERE category_name = 'スポーツ')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '野球');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'バスケットボール', (SELECT category_id FROM categories WHERE category_name = 'スポーツ')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'バスケットボール');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'バレー', (SELECT category_id FROM categories WHERE category_name = 'スポーツ')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'バレー');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'テニス', (SELECT category_id FROM categories WHERE category_name = 'スポーツ')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'テニス');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '国内旅行', (SELECT category_id FROM categories WHERE category_name = 'アウトドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '国内旅行');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '海外旅行', (SELECT category_id FROM categories WHERE category_name = 'アウトドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '海外旅行');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '登山', (SELECT category_id FROM categories WHERE category_name = 'アウトドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '登山');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '釣り', (SELECT category_id FROM categories WHERE category_name = 'アウトドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '釣り');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'キャンプ', (SELECT category_id FROM categories WHERE category_name = 'アウトドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'キャンプ');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '園芸', (SELECT category_id FROM categories WHERE category_name = 'アウトドア')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '園芸');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'アニソン', (SELECT category_id FROM categories WHERE category_name = '音楽')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'アニソン');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'J-POP', (SELECT category_id FROM categories WHERE category_name = '音楽')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'J-POP');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'K-POP', (SELECT category_id FROM categories WHERE category_name = '音楽')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'K-POP');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), '洋楽', (SELECT category_id FROM categories WHERE category_name = '音楽')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = '洋楽');
INSERT INTO hobbies (hobby_id, hobby_name, category_id)
SELECT uuid_generate_v4(), 'クラシック', (SELECT category_id FROM categories WHERE category_name = '音楽')
WHERE NOT EXISTS (SELECT 1 FROM hobbies WHERE hobby_name = 'クラシック');

-- ------------------------------------------
-- matching_questions テーブルへのデフォルトデータ
-- ------------------------------------------
INSERT INTO matching_questions (question_id, question_text)
SELECT uuid_generate_v4(), 'インドア派?' WHERE NOT EXISTS (SELECT 1 FROM matching_questions WHERE question_text = 'インドア派?');
INSERT INTO matching_questions (question_id, question_text)
SELECT uuid_generate_v4(), '国内旅行が好き?' WHERE NOT EXISTS (SELECT 1 FROM matching_questions WHERE question_text = '国内旅行が好き?');
INSERT INTO matching_questions (question_id, question_text)
SELECT uuid_generate_v4(), '競技プログラミングが好き?' WHERE NOT EXISTS (SELECT 1 FROM matching_questions WHERE question_text = '競技プログラミングが好き?');
INSERT INTO matching_questions (question_id, question_text)
SELECT uuid_generate_v4(), 'スポーツが好き?' WHERE NOT EXISTS (SELECT 1 FROM matching_questions WHERE question_text = 'スポーツが好き?');
INSERT INTO matching_questions (question_id, question_text)
SELECT uuid_generate_v4(), 'アニメが趣味?' WHERE NOT EXISTS (SELECT 1 FROM matching_questions WHERE question_text = 'アニメが趣味?');

-- -------------------------------------------
-- question_cards テーブルへのデフォルトデータ
-- -------------------------------------------
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '子供の頃の夢は?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '子供の頃の夢は?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), 'どこに住みたい?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = 'どこに住みたい?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '好きな本は?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '好きな本は?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '最近楽しかったことは?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '最近楽しかったことは?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '密かな楽しみは?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '密かな楽しみは?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '旅行するならどの国?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '旅行するならどの国?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '1番好きな食べ物は?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '1番好きな食べ物は?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), 'もし1億円もらったら何をする?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = 'もし1億円もらったら何をする?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '学生時代の1番の思い出は?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '学生時代の1番の思い出は?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), '出身はどこ?' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = '出身はどこ?');
INSERT INTO question_cards (question_card_id, question_card_text)
SELECT uuid_generate_v4(), 'LINEを交換しましょう!' WHERE NOT EXISTS (SELECT 1 FROM question_cards WHERE question_card_text = 'LINEを交換しましょう!');
