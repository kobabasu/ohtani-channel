# 大谷チャンネル
プロ野球開幕。大谷翔平のプレーを日本で見られるのも、あとわずかかも…  
見逃さないために、打順が回ってきたら、  
自動で中継にチャンネルを合わせる。

## 必要なもの
* IRKit
* GAORAと契約

## install
1. `go build ohtani-channel.go`
1. schedule.jsonに試合日程の日付とスポーツナビの試合番号の対応表を作成
1. IRkitの信号を編集
1. `./ohtani-channel < schedule.json`をcronなどでまわす

## 備考
1. 他の選手でも可。no, team, next, schedule.jsonを編集の上、再コンパイル
1. ただし巨人の亀井など、早打ちの選手だと既に打席が終わってる可能性大
