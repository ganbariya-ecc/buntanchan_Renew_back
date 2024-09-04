async function auth_test() {
    // JWT 取得
    const token = await GetJwt();

    const req = await fetch("/test/atest",{
        method: "POST",
        headers : {
            "Authorized" : token
        }
    })

    console.log(await req.json());
}

function Discord_Auth() {
    window.location.href = "/auth/oauth/discord";
}

function Google_Auth() {
    window.location.href = "/auth/oauth/google";
}

function Line_Auth() {
    window.location.href = "/auth/oauth/line";
}