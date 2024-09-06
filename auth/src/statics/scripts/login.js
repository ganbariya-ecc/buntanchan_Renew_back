// ログインフォーム
const login_form = document.getElementById("login_form");
login_form.addEventListener("submit",async function(evt){
    evt.preventDefault();

    // フォームから値取得
    const username = login_form.elements["username"].value;
    const password = login_form.elements["password"].value;

    // ログイン
    const result = await Login(username,password);

    // 成功したか
    if (result) {
        console.log("ログイン成功");
        window.location.href = "./index.html"
    } else {
        alert("ログイン失敗");
    }
});

async function main() {
    // ログインしているか
    const result = await GetInfo();

    if (result) {
        console.log("ログイン済み");
        window.location.href = "./index.html";
    } else {
        console.log("ログインしてないは");
    }
}

main();