async function Init() {
    // 認証していないときログインに戻す
    await RequireAuth();

    // Admin か オーナーの場合
    if (await IsAdmin() || await IsOwner()) {
        // Adminホームに戻る
        window.location.href = Admin_Home;
    }
}

Init();