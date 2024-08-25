// ログアウトボタン
const logout_btn = document.getElementById("logout_btn");

logout_btn.addEventListener("click", async function (evt) {
    // ログアウト
    const result = await Logout();

    if (result) {
        // 成功した場合
        window.location.href = "./index.html";
    } else {
        // 失敗した場合
        alert("失敗しました");
    }
});

// ユーザー情報のテーブル取得
const users_table = document.getElementById("users_table");

async function main() {
    // ログインしているか
    const result = await GetInfo();

    if (result) {
        console.log("ログイン済み");

        // 初期化
        await init();
    } else {
        console.log("ログインしてないは");
        window.location.href = "./login.html";
    }
}

async function init() {
    const users = await GetUsers();

    users.forEach(user => {
        console.log(user);
        // TR 生成
        const basetr = document.createElement("tr");

        // ユーザーID の th
        const user_idth = document.createElement("th");
        user_idth.scope = "row";
        user_idth.textContent = user["UserID"];
        basetr.appendChild(user_idth);

        // ユーザー名のtd
        const user_nametd = document.createElement("td");
        user_nametd.textContent = user["UserName"];
        basetr.appendChild(user_nametd);

        // 認証方法のtd
        const auth_typetd = document.createElement("td");
        auth_typetd.textContent = user["AuthType"];
        basetr.appendChild(auth_typetd);

        // 認証プロバイダのtd
        const auth_Providertd = document.createElement("td");
        auth_Providertd.textContent = user["Provider"];
        basetr.appendChild(auth_Providertd);

        // ラベルのtd
        const labelstd = document.createElement("td");
        labelstd.textContent = user["Labels"];
        basetr.appendChild(labelstd);

        // 作成日のtd
        const CreatedAt_td = document.createElement("td");
        CreatedAt_td.textContent = user["CreatedAt"];
        basetr.appendChild(CreatedAt_td);

        // 更新日のtd
        const UpdatedAt_td = document.createElement("td");
        UpdatedAt_td.textContent = user["UpdatedAt"];
        basetr.appendChild(UpdatedAt_td);

        users_table.appendChild(basetr);

        basetr.addEventListener("click", function (evt) {
            const userid = user["UserID"];

            window.location.href = `./userinfo.html?userid=${userid}`;
        })
    });
}

main();