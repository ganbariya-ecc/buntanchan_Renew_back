function Init() {
    try {
        const GetUserBtn = document.getElementById("GetUserBtn");

        GetUserBtn.addEventListener("click",async function (evt) {
            console.log(await GetInfo());
        })

        const CreateGroupBtn = document.getElementById("CreateGroupBtn");
        CreateGroupBtn.addEventListener("click",async function (evt) {
            await CreateGrpup("test",members);

            // ページをリロード
            window.location.reload();
        })
    } catch (ex) {
        console.error(ex);
    }
}

const members = [
    {
        "name" : "お父さん",
        "admin" : true,
    },
    {
        "name" : "弟",
        "admin" : false
    },
    {
        "name" : "兄",
        "admin" : true,
    }
]

async function CreateGrpup(name,members) {
    // トークン取得
    const atoken = await GetJwt();

    const req = await fetch("/group/create",{
        method: "POST",
        headers : {
            "Authorization" : atoken,
            "Content-Type" : "application/json",
        },
        body : JSON.stringify({
            "name" : name,
            "members" : members
        })
    })

    console.log(await req.json());
}

// CreateGrpup("test",members);

async function main() {
    try {
        // 現在のグループを取得
        const currentg = await GetCurrentG();

        // グループが存在するとき
        const groupName = currentg["result"]["Group"]["Name"];

        // ロール取得
        const Role = currentg["result"]["Mydata"]["MemberRole"];

        if (groupName != undefined) {
            if (Role == "Owner" || Role == "Admin") {
                window.location.href = Admin_Home;
            } else {
                window.location.href = Member_Home;
            }
        }
    } catch (ex) {
        console.log("エラーです")
        console.error(ex);
    }

    // Body 表示
    document.body.style.display = "";
}

main();

window.addEventListener("DOMContentLoaded",Init)