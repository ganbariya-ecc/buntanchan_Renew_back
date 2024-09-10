const GetUserBtn = document.getElementById("GetUserBtn");

GetUserBtn.addEventListener("click",async function (evt) {
    console.log(await GetInfo());
})

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

const CreateGroupBtn = document.getElementById("CreateGroupBtn");
CreateGroupBtn.addEventListener("click",async function (evt) {
    await CreateGrpup("test",members);
})

// CreateGrpup("test",members);

GetCurrentG();
GetCurrentMembers();