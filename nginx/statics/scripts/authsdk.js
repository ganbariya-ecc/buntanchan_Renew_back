async function Logout() {
    const req = await fetch("/auth/logout",{
        method: "POST",
    });

    if (req.status === 200) {
        console.log(await req.json());
    }
}

async function GetJwt() {
    const req = await fetch("/auth/authed/jwt",{
        method: "POST",
    });

    // 成功したか
    if (req.status === 200) {
        // 成功した場合
        const res = await req.json();
    
        console.log(res["jwt"]);

        // Jwt を返す
        return String(res["jwt"]);
    }

    return ""
}

async function Login(userid,password) {  
    const req = await fetch("/auth/login",{
        method: "POST",
        headers : {
            "Content-Type": "application/json"
        },
        body : JSON.stringify({
            "userid": userid,
            "password": password
        })
    })

    // 成功したか
    if (req.status != 200) {
        console.error(await req.text());
        return 
    }

    console.log(await req.json());
}

async function GetInfo() {
    const req = await fetch("/auth/authed/info",{
        method: "GET",
    });

    // 成功したか
    if (req.status === 200) {
        // 成功した場合
        const res = await req.json();

        // データ を返す
        return res["result"];
    }

    return null;
}

async function RequireAuth() {
    try {
        const myInfo = await GetInfo();

        // ユーザーIDがないとき (ログインしていないとき)
        if (myInfo["UserID"] == "") {
            window.location.href = Login_Page;
        }
    } catch {
        window.location.href = Login_Page;
    }
}