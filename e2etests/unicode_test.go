package e2etests

import (
	_ "embed"
	"testing"
)

func testUnicode(t *testing.T) {
	tcs := []testCase{
		{
			name: "japanese-basic",
			script: `a: ああああああああああ
`,
		},
		{
			name: "japanese-full",
			script: `a: "ある日、トマトが道を歩いていたら、道路の向こうからキュウリがやって来ました。\nトマトは驚いて尋ねました。\n「キュウリさん、どうしてあなたはここにいるのですか？」 キュウリは答えました。「あなたと同じ理由でここにいます。サラダになるために。」"

b: "「バナナは皮を剥いて食べるものです。」" {
  style.font-size: 55
}

a -> b: 「バカは死ななきゃ治らない。」
`,
		},
		{
			name: "emojis",
			script: `a: 🙈🙈🙈🙈🙈🙈🙈🙈
✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊ -> ☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️☁️
`,
		},
		{
			name: "with-style",
			script: `おやすみなさい: {style.stroke-width: 15; style.double-border: true}
`,
		},
		{
			name: "japanese-mixed",
			script: `a: "トマトが赤くなったのはなぜですか？Because it saw the salad dressing!👩‍👩‍👧‍👶👩‍👩‍👧‍👶"
b: "トマトが赤くなったのはなぜですか？Because it saw the salad dressing!👩‍👩‍👧‍👶👩‍👩‍👧‍👶" {
  style.font-size: 100
}
c: 今日はTokyoでsushiを食べました
d: 先日、Shibuyaで友達とshoppingを楽😊しんだ後、ramen屋でdelicious😊なラーメンを食べた。{
  style.font-size: 43
}
e: English English English
f: 先日先日先日
a -> b -> c -> d -> e -> f

`,
		},
	}

	runa(t, tcs)
}
