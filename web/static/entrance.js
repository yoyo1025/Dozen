const CLASSNAME = "bg-visible";
const $target = $(".bg");

// アニメーションを一度だけ実行するように変更
$(document).ready(function () {
    setTimeout(() => {
        $target.addClass(CLASSNAME);
    }, 1500); // ページロード後2秒でアニメーション実行
});
