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

    if (req.status === 200) {
        console.log(await req.json());
    }
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