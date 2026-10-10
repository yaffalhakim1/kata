package main

// Concept is one thing a true beginner has to know before a word list is
// any use: which scripts exist, why a sentence ends in です, what a
// particle does. Each is a short lesson, the way Tae Kim's guide teaches —
// by idea, not by rote list.
type Concept struct {
	Title string
	Body  string
	Example string
	Gloss   string // what the example means, word by word
}

// basics is the first-run lesson: the ideas that make everything else
// click. It is deliberately short and in the order a beginner meets them.
func basics() []Concept {
	return []Concept{
		{
			Title: "Three scripts, one language",
			Body: "Japanese is written with three sets of characters, and they " +
				"do different jobs. Hiragana (あ) spells grammar and native words. " +
				"Katakana (ア) spells foreign words. Kanji (漢字) are Chinese " +
				"characters that carry meaning. You learn hiragana first — " +
				"everything else needs it.",
			Example: "すし　→　寿司　→　スシ",
			Gloss:   "the same word in hiragana, kanji, katakana",
		},
		{
			Title: "There is no word for 'is'",
			Body: "Where English says 'I am a student', Japanese attaches だ " +
				"or its polite form です to the noun: 学生です. There is no " +
				"separate verb 'to be'. To make it negative, swap the ending: " +
				"学生じゃない (is not a student).",
			Example: "学生です。",
			Gloss:   "gakusei desu — '(I) am a student'",
		},
		{
			Title: "Politeness is built in",
			Body: "The same sentence has a casual and a polite form, and the " +
				"polite one is the safe default with strangers and teachers. " +
				"です / ます is polite; dropping them is casual. When unsure, " +
				"be polite.",
			Example: "食べる　／　食べます",
			Gloss:   "taberu (casual) / tabemasu (polite) — 'to eat'",
		},
		{
			Title: "Particles mark the job of a word",
			Body: "Japanese word order is Subject–Object–Verb, and little " +
				"words called particles come after a word to say what job it " +
				"does. は marks the topic, を marks the object, に marks a " +
				"destination or time.",
			Example: "私は 水を 飲みます。",
			Gloss:   "watashi-wa mizu-wo nomimasu — 'I drink water'",
		},
		{
			Title: "The verb comes last",
			Body: "In English the verb comes early: 'I eat sushi.' In Japanese " +
				"it comes at the end: 私は寿司を食べます — literally 'I sushi " +
				"eat'. Once you expect the verb last, sentences stop feeling " +
				"backwards.",
			Example: "私は 寿司を 食べます。",
			Gloss:   "watashi-wa sushi-wo tabemasu — 'I eat sushi'",
		},
		{
			Title: "Long vowels change the word",
			Body: "A vowel held for two beats is a different word. おばさん " +
				"(obasan) is 'aunt'; おばあさん (obaasan) is 'grandmother'. " +
				"In romaji a doubled vowel, or a macron like ō, means 'hold it " +
				"longer'.",
			Example: "おばさん　／　おばあさん",
			Gloss:   "obasan (aunt) / obaasan (grandmother)",
		},
	}
}
