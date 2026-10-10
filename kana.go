package main

// Kana is one hiragana character: how it reads, and a picture to remember
// it by. A learner who cannot read a single kana starts here, because
// nothing else in the language makes sense until they can.
type Kana struct {
	Char     string
	Romaji   string
	Mnemonic string
}

// kanaRow is one column of the gojūon table, which is how kana is taught:
// five vowel sounds, then the same sounds with a consonant in front.
type kanaRow struct {
	Name string // "a", "ka", "sa"…
	Kana []Kana
}

// kanaTable is the hiragana in teaching order. The a-row comes first
// because every other row is those five vowels with a consonant added, so
// getting the vowels right sets up everything after.
func kanaTable() []kanaRow {
	return []kanaRow{
		{"a", []Kana{
			{"あ", "a", "An ant with a big round tummy."},
			{"い", "i", "Two eels, swimming upright."},
			{"う", "u", "A 'u' wearing a little hat."},
			{"え", "e", "An exotic bird with a feather."},
			{"お", "o", "An 'o' doing a handstand."},
		}},
		{"ka", []Kana{
			{"か", "ka", "A mosquito, about to bite."},
			{"き", "ki", "A key, with two teeth."},
			{"く", "ku", "A cuckoo bird's open beak."},
			{"け", "ke", "Wiggly kelp in the sea."},
			{"こ", "ko", "Two cords lying side by side."},
		}},
		{"sa", []Kana{
			{"さ", "sa", "Two hands stirring salsa."},
			{"し", "shi", "A shepherd's crook — sounds like 'she'."},
			{"す", "su", "A swing, loop-de-loop."},
			{"せ", "se", "A mouth with a fang, saying 'se'."},
			{"そ", "so", "A mouth slurping soda."},
		}},
		{"ta", []Kana{
			{"た", "ta", "A fork and a taco."},
			{"ち", "chi", "A mouth saying 'cheese' — 'chi'."},
			{"つ", "tsu", "A great wave, a tsunami."},
			{"て", "te", "A hand-held telescope — 'te' is hand."},
			{"と", "to", "A toe with a splinter."},
		}},
		{"na", []Kana{
			{"な", "na", "A nun praying for nachos."},
			{"に", "ni", "A needle pulling thread."},
			{"ぬ", "nu", "Smooth, bendy noodles."},
			{"ね", "ne", "Nelly the cat, with a curled tail."},
			{"の", "no", "A big pig's nose."},
		}},
		{"ha", []Kana{
			{"は", "ha", "An 'H' and an 'a' — 'ha'!"},
			{"ひ", "hi", "He has a big nose."},
			{"ふ", "fu", "A fool with a hula hoop."},
			{"へ", "he", "A squat little mountain."},
			{"ほ", "ho", "A chimney, with Santa saying 'ho ho'."},
		}},
		{"ma", []Kana{
			{"ま", "ma", "Your mama doing magic."},
			{"み", "mi", "Lucky number twenty-one — 'me'!"},
			{"む", "mu", "A cow saying 'moo'."},
			{"め", "me", "An eye with makeup on — 'me' is eye."},
			{"も", "mo", "More worms on a hook."},
		}},
		{"ya", []Kana{
			{"や", "ya", "A yacht with an anchor."},
			{"ゆ", "yu", "A unique fish looking at you."},
			{"よ", "yo", "A 'Y' and an 'O' — a yo-yo."},
		}},
		{"ra", []Kana{
			{"ら", "ra", "A rabbit with droopy ears."},
			{"り", "ri", "Reeds swaying in the wind."},
			{"る", "ru", "A road with a loop at the end."},
			{"れ", "re", "Someone retching, bent over."},
			{"ろ", "ro", "A plain road, no loop."},
		}},
		{"wa", []Kana{
			{"わ", "wa", "A wasp flying straight up."},
			{"を", "wo", "Marks the object of a sentence. Say 'o'."},
			{"ん", "n", "The only lone consonant, like the n in 'pen'."},
		}},
	}
}

// allKana flattens the table, for the quiz.
func allKana() []Kana {
	var out []Kana
	for _, row := range kanaTable() {
		out = append(out, row.Kana...)
	}
	return out
}
