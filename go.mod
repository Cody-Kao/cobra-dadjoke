module github.com/Cody-Kao/dadjoke
// 以後在cmd就直接呼叫dadjoke就能啟用這個default root command，後面加上其他command名稱就是呼叫他們，例如:joke
// 而dadjoke的由來就是github上的repo name，即:github.com/Cody-Kao/dadjoke

go 1.21.1

require github.com/spf13/cobra v1.8.0

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
