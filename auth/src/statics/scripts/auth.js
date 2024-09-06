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
    console.log(await req.text());

    return true;
}

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
    console.log(await req.text());

    return true;
}

async function Logout() {
    // リクエスト送信
    const req = await fetch("/auth/admin/logout",{
        method: "POST",
    })

    // エラー処理
    if (req.status != 200) {
        // 失敗したとき
        const result = await req.text();
        console.log(`failed to admin logout : ${result}`);
        return false;
    }

    // 結果をjson に
    console.log(await req.json());

    return true;
}

async function GetInfo() {
    // Admin の情報取得
    const req = await fetch("/auth/adminc/info",{
        method: "GET"
    })

    // エラー処理
    if (req.status != 200) {
        const result = await req.text();
        console.log(`failed to get admin info : ${result}`);
        return false;
    }

    const result = await req.text();
    console.log(`admin info : ${result}`);
    return true;
}