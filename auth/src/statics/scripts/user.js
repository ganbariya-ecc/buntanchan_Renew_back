async function GetUsers() {
    const req = await fetch("/auth/adminc/users",{
        method: "GET"
    })

    // エラー処理
    if (req.status != 200) {
        return false;
    }

    const result = await req.json()
    return result["users"];
}