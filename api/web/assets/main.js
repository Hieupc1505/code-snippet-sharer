// function copyCode(element) {
//     let originalHTML = element.innerHTML;
//     element.innerHTML = '<span class="text-green-400 text-sf-s font-bold">Copied</span>';
//     setTimeout(() => {
//         element.innerHTML = originalHTML;
//     }, 3000);
// }


function copyCode(target, element) {
    // Lấy nội dung bên trong thẻ <code> hoặc <pre>
    const textToCopy = target.innerText || target.textContent;

    // Sao chép nội dung vào clipboard
    navigator.clipboard.writeText(textToCopy).then(() => {
        let originalHTML = element.innerHTML; // Lưu nội dung gốc

        // Hiển thị thông báo "Copied"
        element.innerHTML = '<span class="text-green-400 text-sf-s font-bold">Copied</span>';

        // Khôi phục nội dung sau 3 giây
        setTimeout(() => {
            element.innerHTML = originalHTML;
        }, 2000);
    }).catch(err => {
        console.error('Lỗi khi sao chép: ', err);
    });
}



