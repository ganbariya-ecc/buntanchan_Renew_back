const back_button = document.getElementById("back_button");

back_button.addEventListener("click",function(evt){
    window.location.href = "./index.html"
})

// URL オブジェクト
let url_string = window.location.href;
// 文字列としてのURLをURLオブジェクトに変換する。
let url = new URL(url_string);
// URLオブジェクトのsearchParamsのget関数でIDがuseridの値を取得する。
let userid = url.searchParams.get("userid");

// タグ保存ボタン
const tags_save_button = document.getElementById("tags_save_button");

tags_save_button.addEventListener("click",async function (evt) {
    // タグ一覧取得
    const tags = test_input.value.split(",");

    // リクエスト
    const req = await fetch("/auth/adminc/labels/update",{
        method: "POST",
        headers : {
            "Content-Type" : "application/json",
            "userid" : userid,
        },
        body : JSON.stringify({
            "labels" : tags
        })
    });

    // エラー処理
    if (req.status != 200) {
        console.error(await req.text())
        return
    }

    const result = await req.json()

    console.log(result);
})

// ユーザー削除ボタン
const delete_btn = document.getElementById("delete_btn");

delete_btn.addEventListener("click",async function (evt) {
    // 確認ダイアログ
    const check = confirm("ユーザーを削除しますか?");

    // 確認
    if (!check) {
        console.log("戻る");
        return;
    }

    // リクエスト
    const req = await fetch("/auth/adminc/user/delete",{
        method : "DELETE",
        headers : {
            "userid" : userid
        }
    })

    // 成功したか
    if (req.status != 200) {
        console.error(await req.text());
        alert("削除失敗");
        return;
    }

    // 成功したとき
    console.log(await req.text())

    // index に戻る
    window.location.href = "/auth/index.html";
})

async function main() {
    const test_input = document.getElementById('test_input');

    UseBootstrapTag(test_input);

    test_input.addEventListener("change",function(evt){
        console.log(test_input.value.split(","));
    })

    console.log(userid);

    // ユーザー情報取得
    const req = await fetch("/auth/adminc/user/info",{
        method: "GET",
        headers : {
            "userid" : userid
        }
    })

    // エラー処理
    if (req.status != 200) {
        console.error(await req.text());
            
        // index に戻る
        window.location.href = "/auth/index.html";
        return;
    }

    // ユーザー情報表示
    console.log(await req.json());
}

main();