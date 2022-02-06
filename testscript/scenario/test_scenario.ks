/** @lisence KScript2 System ⒸKoromosoft 2021.
* @ks_version 2.0
* script ch1 000
* @writer Chicchi
* @scripter Chicchi
*/
import "alice2_common.ks"

// 条件：ゲーム開始時に自動開始
func scene_prolog1()
func scene_prolog2()

func main(){
    Alice2_Init()
    
    BB_In()

    // シナリオ開始
    scene_prolog1()
    scene_prolog2()

    BB_Out()

    Alice2_Shut()
}

func scene_prolog1(){
    Print(P_STILL_001)
    スチル000 := NewImg(P_STILL_000)
    スチル001 := NewImg(P_STILL_001)
    スチル002 := NewImg(P_STILL_002)
    フィロ := NewImg(P_ST_001_Firo)
    フィロ拡大 := NewImg(P_ST_001_Firo)
    車掌 := NewImg(P_ST_007_Wreath)
    背景000 := NewImg(P_BG_000)
    背景000拡大 := NewImg(P_BG_000)

    #novel{
        //明点
        @WhiteOut

        @ShowWindow
        - nil
        ひらり、と、真っ赤な花びらが、
        車窓から舞い降りてくる。」
        差し込んでくる西日を手で遮りながら外を見ると、
        いち面の赤い花が見えた。」
        そんなぼくの様子を見ていた”彼女”が、
        おもむろに立ち上がって、車窓から外を覗く。」
        // スチル：000 列車の中のポーレット
        @Still スチル000 0 3
        @ClrFilter_3s

        - 女の子
        わあ！　すごいよ！
        一面、”魔女の花”だらけだ。」
        - nil
        汽車はいつの間にか、花畑の中を走っていた。」
        - 女の子
        ねえルイス。
        わたしたち、これからどこへ行くのかな。」
        - nil
        彼女はポケットから小さな紙切れを取り出す。」
        それは、読めない文字が印刷された切符。」
        - ルイス
        ぼくのと、色が違うんだね。」
        - nil

        // スチル：001 二枚のきっぷ　オーバーレイ
        @Still スチル001 0 4
        - nil
        ぼくのは緑色。彼女のは赤色。」
        - 女の子
        ルイスは男の子だから緑。
        わたしは、女の子だから赤色。」
        きっとそういうことだよ！」
        - nil
        男女で切符の色が違うなんてきいたことない。
        そう言おうとしたとき――」
        // 明点
        @WhiteFade
        @ClrAsync スチル000
        @ClrAsync スチル001
        @Await

        - nil
        大きな汽笛の音とともに、
        ”世界が崩れた”。」
        // 明点クリア
        // 背景：000 列車の中
        // 立ち絵表示：ポーレット　拡大
        @BGAsync 背景000 0
        @Face フィロ 0 "通常" "通常"
        @STAsync フィロ -1 1
        @Await
        @ClrFilter
        - 女の子
        みて……花が……。」
        - nil
        いままで広がっていた景色は、
        細かい粒子になって、登っていく。」
        - 女の子
        あの花びらはどこに行くのかな。
        あんなに高く登って……。」
        - ？？？
        命になるんだよ。」
        - nil

        // 立ち絵表示：ポーレット
        // 立ち絵表示：車掌
        @Face 車掌 0 "通常" "通常"
        @STAsync 車掌 1 1
        @Await
        - nil
        いつのまにか、すぐ横の通路に
        車掌服を着た”なにか”が立っていた。」
        帽子をふかくかぶり、体に対して、少し大きめの
        ジャケットの襟が覆っているので、顔が見えない。」
        声も中性的で、男なのか女なのか、
        子供なのか老人なのかすら判別できない、」
        不思議な”なにか”だった。」
        - ルイス
        命になるって、
        それはどういう例えで……。」
        - 車掌
        文字通りの意味さ。」
        ぼうや。
        ここから先は”逢魔の世界”だ。」
        君たちの世界とは、なにもかもが違う。」
        だから、花びらが集まって、一つの生命になる……
        なんてことも、この世界では当たり前なんだよ。」
        ……おっと、あっちの女の子は違うようだけれど、
        君は”見送り”だね。」
        なら、そろそろ降りなきゃいけない。」
        - ルイス
        何を言って……。」
        - 車掌
        切符の色だよ。」
        君は、あの子とは行き先が違う。
        だから、君はここまで……だ。」
        - ルイス
        だから、意味がわからな……。」
        - nil
        そっと、車掌がぼくの頭に手を乗せる。」
        - 車掌
        大丈夫。
        彼女とはきっと、また会えるから。」
        その願いを持ち続けている限り、
        ”逢魔の世界”は、君の願いを叶えてくれるだろう。」
        - ルイス
        また会える……。」
        - nil
        とたんにぼくは、
        とても大切なことを思い出した。」
        彼女を……助けなきゃ。」
        そう、目の前にいる彼女は、
        まだ”救われていない”のだから。」
        // 立ち絵表示：ポーレット　拡大
        @BGZoomAsync 背景000拡大 0
        @ClrAsync 背景000
        @Face フィロ拡大 0 "通常" "通常"
        @STZoomAsync フィロ拡大 1
        @ClrAsync フィロ
        @ClrAsync 車掌
        @Await
        ぼくは、彼女の方を見る。」
        @Face フィロ拡大 0 "通常" "笑顔"
        彼女は『なあに？』というような顔で、
        こちらをみる。」
        - ルイス
        かならず、助けるから。
        ぼくが必ず……。」
        - nil

        //汽笛の音
        //明点
        @WhiteOut
        @ClrAsync フィロ拡大
        @ClrAsync 背景000拡大
        @Await
        - ルイス
        ―――…――……！！」
        - nil
        彼女の名前を叫ぶ。」
        つらいことも、悲しいことも、君の身に起こる
        全部の悪いことから、ぼくが守ってみせる。」
        そんな誓いも、
        今は汽笛の音にかき消された。」

        @HideWindow

        // スチル　002 タイトル
        @Still スチル002 0 3
        @ClrFilter_3s
        」
        @ClrAsync スチル002
        @Await
        
        // 明点　シーン終了
    }

    DeleteImg(スチル000)
    DeleteImg(スチル001)
    DeleteImg(スチル002)
    DeleteImg(フィロ)
    DeleteImg(フィロ拡大)
    DeleteImg(車掌)
    DeleteImg(背景000)
    DeleteImg(背景000拡大)
}

func scene_prolog2(){
    スチル003 := NewImg("Still/still_003")
    
    #novel{
        @BlackFade
        @ShowWindow

        - ？？？
        …ーーー。　…ーー…！」
        - ルイス
        う……ん。」
        - アリス
        ……ーい、おいったら！」
        - nil
        
        // イベントスチル　アリスと汽車の中
        @Still スチル003 1 1
        @ClrFilter
        - ルイス
        ああ、ごめん。
        寝てたかも……。」
        - アリス
        のんきなもんだなあ。
        明日からいよいよ”女王選”が始まるってのに。」
        - ルイス
        そうだね。ちょっと疲れていたのかも。
        気合を入れ直すよ！」
        - アリス
        クフフ！　いい心がけだぞ。」
        この私がプリンセスを引き受けたんだ。
        絶対に勝ち残ってやるから、任せておけ！」
        - ルイス
        アキタニア公がわざわざ国外から呼んだんだ。」
        思ったより普通の女の子だと思ったけど、
        自信はありそうだね。」
        - アリス
        普通の女の子は余計だろ？」
        お前こそ、”後見人”なんだから、
        ちゃんとついてこいよな。」
        私は正直不安だぞ？」
        お前みたいな、ひょろひょろなガキに
        背中を預けなきゃいけないんだ。」
        - ルイス
        だ、誰がガキだ！
        これでも20歳だぞ。」
        この国では立派なオトナだ。
        ……まあ、背は小さいけどさ。」
        - アリス
        お前も後見人になったからには、
        地位とか金がほしいんだろ？」
        しっかり働かないと、
        女王権限でクビにしちまうぞ？」
        - ルイス
        ぼくは別に……、
        そんなものに興味はないよ。」
        - アリス
        あん？　じゃあなんでわざわざ
        こんな面倒なことを引き受けたんだ？」
        - ルイス
        ……ある人を助けるためだ。」
        - アリス
        ある人……、彼女か？
        ウブな顔して、意外とやるんだな。」
        - ルイス
        そ、そんなんじゃない！」
        妹だよ……。」
        - nil

        @ChangeStill スチル003 0
        - nil
        汽車は山間を進み、トンネルを抜け、
        遠くに町が見えてきた。」
        - アリス
        あれが私達の本拠か！」
        - ルイス
        うん。
        僕らの国、アキタニア公国の首都ボルドンだ。」
        - アリス
        クフフッ！
        なんかわくわくしてきたな！」
        おいルイス。この女王選、
        絶対に勝ち残るぞ！」
        事情はよく知らんが、
        お前の妹も助けて、それでハッピーエンドだ。」
        - ルイス
        そうだね……うん。」
        アリス、ありがとう。
        改めてよろしくね。」
        - アリス
        おう！　頼りにしてるぞ、相棒。」
        - nil

        @WhiteOut
        @ClrAsync スチル003
        @Await
        - nil
        窓から風が吹き込んで、
        赤い花びらがひらり……、僕らの手の中に舞い降りた。」
        - ルイス
        魔女の花……。」
        - nil
        汽車はボルドンの手前の平野に差し掛かっていた。」
        平野には、一面魔女の花が咲き誇っている。」
        まるで、僕らの出発を祝福するかのように。」
        - ルイス
        ……必ず助けるからね。」
        ポーレット……。」
        - nil
        ぼくらのながい、ながい戦いが、ここからはじまる。」
        @ClrFilter
        // 汽笛の音。明点　オープニング再生

        @HideWindow
    }

    DeleteImg(スチル003)
}
