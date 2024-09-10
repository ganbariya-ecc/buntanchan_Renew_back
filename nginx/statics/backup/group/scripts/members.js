const members_area = document.getElementById("members_area");

async function Init() {
    // 認証必須にする
    await RequireAuth();

    // admin かオーナー以外戻す
    if (await IsAdmin() || await IsOwner()) {
        const members = await GetCurrentMembers();

        members.forEach(member => {
            const userdiv = document
        });
    }
}

Init();