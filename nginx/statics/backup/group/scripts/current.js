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

    const result =  await req.json();

    return result;
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

    const result = await req.json();

    return result["result"];
}

async function IsOwner() {
    // 現在のグループを取得
    const currentg = await GetCurrentG();
    // ロール取得
    const Role = currentg["result"]["Mydata"]["MemberRole"];

    return Role == "Owner";
}

async function IsAdmin() {
    // 現在のグループを取得
    const currentg = await GetCurrentG();
    // ロール取得
    const Role = currentg["result"]["Mydata"]["MemberRole"];

    return Role == "Admin";
}

async function IsMember() {
    // 現在のグループを取得
    const currentg = await GetCurrentG();
    // ロール取得
    const Role = currentg["result"]["Mydata"]["MemberRole"];

    return Role == "Member";
}