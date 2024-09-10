async function Init() {
    // 認証していないときログインに戻す
    await RequireAuth();

    // Admin か オーナーじゃない場合
    if (await IsAdmin() || await IsOwner()) {
    } else {
        // メンバーホームに戻る
        window.location.href = Member_Home;
    }
}

Init();