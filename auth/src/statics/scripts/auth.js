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
    } else {
        console.log("ログイン失敗");
    }
});

// ログイン
async function Login(username,password) {
    // ログイン
    const req = await fetch("/auth/admin/login",{
        method: "POST",
        headers : {
            "Content-Type" : "application/json"
        },
        body : JSON.stringify({
            "username" : username,
            "password" : password
        })
    })

    // エラー処理
    if (req.status != 200) {
        // 失敗したとき
        console.log(`failed to admin login : ${await req.text()}`);
        return false;
    }

    // 結果をjson に
    console.log(await req.json());

    return true;
}

// サインアップフォーム
const signup_form = document.getElementById("signup_form");
signup_form.addEventListener("submit",async function(evt){
    evt.preventDefault();

    // フォームから値取得
    const username = login_form.elements["username"].value;
    const password = login_form.elements["password"].value;

    // ログイン
    const result = await Signup(username,password);

    // 成功したか
    if (result) {
        console.log("サインアップ成功");
    } else {
        console.log("サインアップ失敗");
    }
});

// サインアップ
async function Signup(username,password) {
    // サインアップ
    const req = await fetch("/auth/admin/signup",{
        method: "POST",
        headers : {
            "Content-Type" : "application/json"
        },
        body : JSON.stringify({
            "username" : username,
            "password" : password
        })
    })

    // エラー処理
    if (req.status != 200) {
        // 失敗したとき
        console.log(`failed to admin signup : ${await req.text()}`);
        return false;
    }

    // 結果をjson に
    console.log(await req.json());

    return true;
}