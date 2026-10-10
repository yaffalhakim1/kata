package main

import "strings"

// Card is one thing to learn: a Japanese word, how to read it, its
// Indonesian meaning, and a short Japanese example so the word is seen in
// use.
type Card struct {
	Japanese string // the word, in kana or kanji
	Romaji   string // how it is read, for a beginner
	Indo     string // the Indonesian meaning
	Example  string // a short example sentence
	Topic    string
}

// deck is the built-in word list. It is small on purpose: a learner needs
// common words first, and a list that fits in one file stays correct.
func deck() []Card {
	return []Card{
		// Salam — greetings.
		{Japanese: "こんにちは", Romaji: "konnichiwa", Indo: "halo", Example: "こんにちは、元気ですか？", Topic: "Salam"},
		{Japanese: "おはよう", Romaji: "ohayou", Indo: "selamat pagi", Example: "おはよう、先生。", Topic: "Salam"},
		{Japanese: "こんばんは", Romaji: "konbanwa", Indo: "selamat malam", Example: "こんばんは、また明日。", Topic: "Salam"},
		{Japanese: "ありがとう", Romaji: "arigatou", Indo: "terima kasih", Example: "ありがとう、助かりました。", Topic: "Salam"},
		{Japanese: "どういたしまして", Romaji: "dou itashimashite", Indo: "sama-sama", Example: "どういたしまして。", Topic: "Salam"},
		{Japanese: "すみません", Romaji: "sumimasen", Indo: "maaf / permisi", Example: "すみません、遅れました。", Topic: "Salam"},
		{Japanese: "ごめんなさい", Romaji: "gomen nasai", Indo: "maaf", Example: "ごめんなさい、私のミスです。", Topic: "Salam"},
		{Japanese: "さようなら", Romaji: "sayounara", Indo: "selamat tinggal", Example: "さようなら、また明日。", Topic: "Salam"},

		// Makanan — food.
		{Japanese: "水", Romaji: "mizu", Indo: "air", Example: "水を飲みます。", Topic: "Makanan"},
		{Japanese: "ご飯", Romaji: "gohan", Indo: "nasi", Example: "ご飯を食べます。", Topic: "Makanan"},
		{Japanese: "鶏肉", Romaji: "toriniku", Indo: "ayam", Example: "鶏肉が好きです。", Topic: "Makanan"},
		{Japanese: "卵", Romaji: "tamago", Indo: "telur", Example: "卵を二つください。", Topic: "Makanan"},
		{Japanese: "パン", Romaji: "pan", Indo: "roti", Example: "パンを買います。", Topic: "Makanan"},
		{Japanese: "コーヒー", Romaji: "koohii", Indo: "kopi", Example: "朝、コーヒーを飲みます。", Topic: "Makanan"},
		{Japanese: "辛い", Romaji: "karai", Indo: "pedas", Example: "この料理は辛いです。", Topic: "Makanan"},
		{Japanese: "甘い", Romaji: "amai", Indo: "manis", Example: "このケーキは甘いです。", Topic: "Makanan"},
		{Japanese: "食べる", Romaji: "taberu", Indo: "makan", Example: "一緒に食べましょう。", Topic: "Makanan"},
		{Japanese: "飲む", Romaji: "nomu", Indo: "minum", Example: "水をもっと飲んでください。", Topic: "Makanan"},

		// Sehari-hari — daily life.
		{Japanese: "家", Romaji: "ie", Indo: "rumah", Example: "私の家は駅の近くです。", Topic: "Sehari-hari"},
		{Japanese: "友達", Romaji: "tomodachi", Indo: "teman", Example: "彼女は私の友達です。", Topic: "Sehari-hari"},
		{Japanese: "家族", Romaji: "kazoku", Indo: "keluarga", Example: "家族はここに住んでいます。", Topic: "Sehari-hari"},
		{Japanese: "仕事", Romaji: "shigoto", Indo: "pekerjaan", Example: "八時に仕事に行きます。", Topic: "Sehari-hari"},
		{Japanese: "お金", Romaji: "okane", Indo: "uang", Example: "今日はお金がありません。", Topic: "Sehari-hari"},
		{Japanese: "時間", Romaji: "jikan", Indo: "waktu", Example: "今、何時ですか？", Topic: "Sehari-hari"},
		{Japanese: "日", Romaji: "hi", Indo: "hari", Example: "今日はいい日です。", Topic: "Sehari-hari"},
		{Japanese: "夜", Romaji: "yoru", Indo: "malam", Example: "夜は静かです。", Topic: "Sehari-hari"},
		{Japanese: "今日", Romaji: "kyou", Indo: "hari ini", Example: "今日、日本語を勉強します。", Topic: "Sehari-hari"},
		{Japanese: "明日", Romaji: "ashita", Indo: "besok", Example: "また明日。", Topic: "Sehari-hari"},

		// Sifat — describing.
		{Japanese: "大きい", Romaji: "ookii", Indo: "besar", Example: "それは大きい町です。", Topic: "Sifat"},
		{Japanese: "小さい", Romaji: "chiisai", Indo: "kecil", Example: "小さい部屋です。", Topic: "Sifat"},
		{Japanese: "いい", Romaji: "ii", Indo: "bagus", Example: "これはいい考えです。", Topic: "Sifat"},
		{Japanese: "悪い", Romaji: "warui", Indo: "buruk", Example: "天気が悪いです。", Topic: "Sifat"},
		{Japanese: "暑い", Romaji: "atsui", Indo: "panas", Example: "今日は暑いです。", Topic: "Sifat"},
		{Japanese: "寒い", Romaji: "samui", Indo: "dingin", Example: "冬は寒いです。", Topic: "Sifat"},
		{Japanese: "速い", Romaji: "hayai", Indo: "cepat", Example: "この電車は速いです。", Topic: "Sifat"},
		{Japanese: "遅い", Romaji: "osoi", Indo: "lambat", Example: "今日はバスが遅いです。", Topic: "Sifat"},
		{Japanese: "嬉しい", Romaji: "ureshii", Indo: "senang", Example: "会えて嬉しいです。", Topic: "Sifat"},
		{Japanese: "疲れた", Romaji: "tsukareta", Indo: "lelah", Example: "仕事の後で疲れました。", Topic: "Sifat"},

		// Kata Kerja — verbs.
		{Japanese: "行く", Romaji: "iku", Indo: "pergi", Example: "バスで学校に行きます。", Topic: "Kata Kerja"},
		{Japanese: "来る", Romaji: "kuru", Indo: "datang", Example: "私の家に来てください。", Topic: "Kata Kerja"},
		{Japanese: "見る", Romaji: "miru", Indo: "melihat", Example: "鳥を見ます。", Topic: "Kata Kerja"},
		{Japanese: "話す", Romaji: "hanasu", Indo: "berbicara", Example: "日本語を話しますか？", Topic: "Kata Kerja"},
		{Japanese: "聞く", Romaji: "kiku", Indo: "mendengar", Example: "先生の話を聞いてください。", Topic: "Kata Kerja"},
		{Japanese: "読む", Romaji: "yomu", Indo: "membaca", Example: "毎晩、本を読みます。", Topic: "Kata Kerja"},
		{Japanese: "書く", Romaji: "kaku", Indo: "menulis", Example: "手紙を書きます。", Topic: "Kata Kerja"},
		{Japanese: "買う", Romaji: "kau", Indo: "membeli", Example: "市場で果物を買います。", Topic: "Kata Kerja"},
		{Japanese: "助ける", Romaji: "tasukeru", Indo: "menolong", Example: "助けてください。", Topic: "Kata Kerja"},
		{Japanese: "勉強する", Romaji: "benkyou suru", Indo: "belajar", Example: "日本語を勉強します。", Topic: "Kata Kerja"},
	}
}

// topicNames lists the deck's topics in a fixed order, for the filter.
func topicNames(cards []Card) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, c := range cards {
		if !seen[c.Topic] {
			seen[c.Topic] = true
			out = append(out, c.Topic)
		}
	}
	return out
}

// cardsIn returns the cards of one topic, or the whole deck for "".
func cardsIn(cards []Card, topic string) []Card {
	if topic == "" {
		return cards
	}
	out := make([]Card, 0, len(cards))
	for _, c := range cards {
		if c.Topic == topic {
			out = append(out, c)
		}
	}
	return out
}

// filterCards keeps the cards matching a query, which matches the Japanese,
// the reading, the meaning and the topic, case-insensitively.
func filterCards(cards []Card, query string) []Card {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return cards
	}
	out := make([]Card, 0, len(cards))
	for _, c := range cards {
		if strings.Contains(strings.ToLower(c.Japanese), q) ||
			strings.Contains(strings.ToLower(c.Romaji), q) ||
			strings.Contains(strings.ToLower(c.Indo), q) ||
			strings.Contains(strings.ToLower(c.Topic), q) {
			out = append(out, c)
		}
	}
	return out
}
