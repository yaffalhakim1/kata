package main

import "strings"

// Card is one thing to learn: a Japanese word, how to read it, its
// Indonesian meaning, a short example, and a note that teaches how the
// word is actually used — because a beginner needs the idea, not just the
// gloss.
type Card struct {
	Japanese string // the word, in kana or kanji
	Romaji   string // how it is read, for a beginner
	Indo     string // the Indonesian meaning
	Example  string // a short example sentence
	Note     string // how it is used, in plain Indonesian
	Topic    string
}

// deck is the built-in word list. It is small on purpose: a learner needs
// common words first, and a list that fits in one file stays correct.
func deck() []Card {
	return []Card{
		// Salam — greetings.
		{Japanese: "こんにちは", Romaji: "konnichiwa", Indo: "halo", Example: "こんにちは、元気ですか？", Note: "Dipakai siang hari. Pagi おはよう, malam こんばんは.", Topic: "Salam"},
		{Japanese: "おはよう", Romaji: "ohayou", Indo: "selamat pagi", Example: "おはよう、先生。", Note: "Bentuk sopan: おはようございます.", Topic: "Salam"},
		{Japanese: "こんばんは", Romaji: "konbanwa", Indo: "selamat malam", Example: "こんばんは、また明日。", Note: "Dipakai saat bertemu malam hari, bukan saat berpisah.", Topic: "Salam"},
		{Japanese: "ありがとう", Romaji: "arigatou", Indo: "terima kasih", Example: "ありがとう、助かりました。", Note: "Sopan: ありがとうございます. Jangan pakai ke atasan.", Topic: "Salam"},
		{Japanese: "どういたしまして", Romaji: "dou itashimashite", Indo: "sama-sama", Example: "どういたしまして。", Note: "Jawaban dari ありがとう.", Topic: "Salam"},
		{Japanese: "すみません", Romaji: "sumimasen", Indo: "maaf / permisi", Example: "すみません、遅れました。", Note: "Serbaguna: permisi, minta maaf, memanggil pelayan.", Topic: "Salam"},
		{Japanese: "ごめんなさい", Romaji: "gomen nasai", Indo: "maaf", Example: "ごめんなさい、私のミスです。", Note: "Untuk kesalahan sungguhan. Bentuk santai: ごめん.", Topic: "Salam"},
		{Japanese: "さようなら", Romaji: "sayounara", Indo: "selamat tinggal", Example: "さようなら、また明日。", Note: "Untuk perpisahan lama. Sehari-hari cukup じゃあね.", Topic: "Salam"},

		// Makanan — food.
		{Japanese: "水", Romaji: "mizu", Indo: "air", Example: "水を飲みます。", Note: "Benda, jadi pakai partikel を sebagai objek.", Topic: "Makanan"},
		{Japanese: "ご飯", Romaji: "gohan", Indo: "nasi", Example: "ご飯を食べます。", Note: "ご飯 juga berarti 'makan' itu sendiri.", Topic: "Makanan"},
		{Japanese: "鶏肉", Romaji: "toriniku", Indo: "ayam", Example: "鶏肉が好きです。", Note: "Suka sesuatu pakai が, bukan を.", Topic: "Makanan"},
		{Japanese: "卵", Romaji: "tamago", Indo: "telur", Example: "卵を二つください。", Note: "Minta sesuatu: kata benda + をください.", Topic: "Makanan"},
		{Japanese: "パン", Romaji: "pan", Indo: "roti", Example: "パンを買います。", Note: "Ditulis katakana karena kata serapan (dari Portugis).", Topic: "Makanan"},
		{Japanese: "コーヒー", Romaji: "koohii", Indo: "kopi", Example: "朝、コーヒーを飲みます。", Note: "Tanda ー berarti vokal panjang.", Topic: "Makanan"},
		{Japanese: "辛い", Romaji: "karai", Indo: "pedas", Example: "この料理は辛いです。", Note: "Kata sifat-i. 辛いです = 'pedas'.", Topic: "Makanan"},
		{Japanese: "甘い", Romaji: "amai", Indo: "manis", Example: "このケーキは甘いです。", Note: "Kata sifat berakhiran い, berubah jadi くない untuk negatif.", Topic: "Makanan"},
		{Japanese: "食べる", Romaji: "taberu", Indo: "makan", Example: "一緒に食べましょう。", Note: "Sopan: 食べます. Ayo: 食べましょう.", Topic: "Makanan"},
		{Japanese: "飲む", Romaji: "nomu", Indo: "minum", Example: "水をもっと飲んでください。", Note: "Minta tolong: 〜てください.", Topic: "Makanan"},

		// Sehari-hari — daily life.
		{Japanese: "家", Romaji: "ie", Indo: "rumah", Example: "私の家は駅の近くです。", Note: "私の = 'rumah saya'. の menyatakan kepunyaan.", Topic: "Sehari-hari"},
		{Japanese: "友達", Romaji: "tomodachi", Indo: "teman", Example: "彼女は私の友達です。", Note: "Sopan: 友達です. Santai cukup 友達.", Topic: "Sehari-hari"},
		{Japanese: "家族", Romaji: "kazoku", Indo: "keluarga", Example: "家族はここに住んでいます。", Note: "Sedang berlangsung: 〜ています.", Topic: "Sehari-hari"},
		{Japanese: "仕事", Romaji: "shigoto", Indo: "pekerjaan", Example: "八時に仕事に行きます。", Note: "Pergi ke tempat pakai partikel に.", Topic: "Sehari-hari"},
		{Japanese: "お金", Romaji: "okane", Indo: "uang", Example: "今日はお金がありません。", Note: "Tidak punya: がありません.", Topic: "Sehari-hari"},
		{Japanese: "時間", Romaji: "jikan", Indo: "waktu", Example: "今、何時ですか？", Note: "Bertanya jam: 何時ですか.", Topic: "Sehari-hari"},
		{Japanese: "日", Romaji: "hi", Indo: "hari", Example: "今日はいい日です。", Note: "Kanji ini berubah baca: 日 hi, 日本 nihon.", Topic: "Sehari-hari"},
		{Japanese: "夜", Romaji: "yoru", Indo: "malam", Example: "夜は静かです。", Note: "Kata benda + です = 'adalah ...'.", Topic: "Sehari-hari"},
		{Japanese: "今日", Romaji: "kyou", Indo: "hari ini", Example: "今日、日本語を勉強します。", Note: "Waktu tidak butuh partikel; bisa di depan kalimat.", Topic: "Sehari-hari"},
		{Japanese: "明日", Romaji: "ashita", Indo: "besok", Example: "また明日。", Note: "また明日 = 'sampai besok'.", Topic: "Sehari-hari"},

		// Sifat — describing.
		{Japanese: "大きい", Romaji: "ookii", Indo: "besar", Example: "それは大きい町です。", Note: "Kata sifat-i menempel langsung ke kata benda.", Topic: "Sifat"},
		{Japanese: "小さい", Romaji: "chiisai", Indo: "kecil", Example: "小さい部屋です。", Note: "Antonim dari 大きい.", Topic: "Sifat"},
		{Japanese: "いい", Romaji: "ii", Indo: "bagus", Example: "これはいい考えです。", Note: "Negatifnya tidak beraturan: いい → よくない.", Topic: "Sifat"},
		{Japanese: "悪い", Romaji: "warui", Indo: "buruk", Example: "天気が悪いです。", Note: "Antonim dari いい.", Topic: "Sifat"},
		{Japanese: "暑い", Romaji: "atsui", Indo: "panas", Example: "今日は暑いです。", Note: "暑い untuk cuaca; 熱い untuk benda panas.", Topic: "Sifat"},
		{Japanese: "寒い", Romaji: "samui", Indo: "dingin", Example: "冬は寒いです。", Note: "寒い untuk udara; 冷たい untuk benda dingin.", Topic: "Sifat"},
		{Japanese: "速い", Romaji: "hayai", Indo: "cepat", Example: "この電車は速いです。", Note: "Berhati-hati: 早い juga 'hayai' tapi artinya 'awal'.", Topic: "Sifat"},
		{Japanese: "遅い", Romaji: "osoi", Indo: "lambat", Example: "今日はバスが遅いです。", Note: "Antonim dari 速い.", Topic: "Sifat"},
		{Japanese: "嬉しい", Romaji: "ureshii", Indo: "senang", Example: "会えて嬉しいです。", Note: "Perasaan orang pertama, bukan orang lain.", Topic: "Sifat"},
		{Japanese: "疲れた", Romaji: "tsukareta", Indo: "lelah", Example: "仕事の後で疲れました。", Note: "Bentuk lampau dari 疲れる, jadi 'sudah lelah'.", Topic: "Sifat"},

		// Kata Kerja — verbs.
		{Japanese: "行く", Romaji: "iku", Indo: "pergi", Example: "バスで学校に行きます。", Note: "Pergi ke suatu tempat pakai に. Sopan: 行きます.", Topic: "Kata Kerja"},
		{Japanese: "来る", Romaji: "kuru", Indo: "datang", Example: "私の家に来てください。", Note: "Kata kerja tidak beraturan; sopan: 来ます.", Topic: "Kata Kerja"},
		{Japanese: "見る", Romaji: "miru", Indo: "melihat", Example: "鳥を見ます。", Note: "Yang dilihat pakai partikel を.", Topic: "Kata Kerja"},
		{Japanese: "話す", Romaji: "hanasu", Indo: "berbicara", Example: "日本語を話しますか？", Note: "Bertanya: tambahkan か di akhir.", Topic: "Kata Kerja"},
		{Japanese: "聞く", Romaji: "kiku", Indo: "mendengar / bertanya", Example: "先生の話を聞いてください。", Note: "Satu kata, dua arti: mendengar dan bertanya.", Topic: "Kata Kerja"},
		{Japanese: "読む", Romaji: "yomu", Indo: "membaca", Example: "毎晩、本を読みます。", Note: "Kebiasaan: pakai bentuk 〜ます.", Topic: "Kata Kerja"},
		{Japanese: "書く", Romaji: "kaku", Indo: "menulis", Example: "手紙を書きます。", Note: "Objek (手紙) ditandai partikel を.", Topic: "Kata Kerja"},
		{Japanese: "買う", Romaji: "kau", Indo: "membeli", Example: "市場で果物を買います。", Note: "Tempat tindakan pakai で, bukan に.", Topic: "Kata Kerja"},
		{Japanese: "助ける", Romaji: "tasukeru", Indo: "menolong", Example: "助けてください。", Note: "助けて！ = 'tolong!' saat darurat.", Topic: "Kata Kerja"},
		{Japanese: "勉強する", Romaji: "benkyou suru", Indo: "belajar", Example: "日本語を勉強します。", Note: "Kata benda + する menjadi kata kerja.", Topic: "Kata Kerja"},
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
