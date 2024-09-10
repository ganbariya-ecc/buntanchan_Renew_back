// 現在のグループ取得
async function GetCurrentG() {
    const atoken = await GetJwt();

    const req = await fetch("/group/current",{
        method: "GET",
        headers : {
            "Authorization" : atoken,
            "Content-Type" : "application/json",
        },
    })

    console.log(await req.json());
}

// 現在のメンバー取得
async function GetCurrentMembers() {
    const atoken = await GetJwt();

    const req = await fetch("/group/current/members",{
        method: "GET",
        headers : {
            "Authorization" : atoken,
            "Content-Type" : "application/json",
        },
    })

    console.log(await req.json());
}