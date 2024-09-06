// サインアップフォーム
const signup_form = document.getElementById("signup_form");
signup_form.addEventListener("submit", async function (evt) {
    evt.preventDefault();

    // フォームから値取得
    const username = signup_form.elements["username"].value;
    const password = signup_form.elements["password"].value;
    const confirm_password = signup_form.elements["confirm_password"].value;

    // パスワード検証
    if (password != confirm_password) {
        alert("パスワードが一致しません");
        return;
    }

    // ログイン
    const result = await Signup(username, password);

    // 成功したか
    if (result) {
        window.location.href = "./index.html";
        console.log("サインアップ成功");
    } else {
        alert("サインアップ失敗");
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