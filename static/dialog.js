const dialog = (function(){
    const exports = {};

    const messageDialog = document.createElement('dialog');
    document.body.appendChild(messageDialog);
    const messageP = document.createElement('p');
    messageDialog.appendChild(messageP);
    const messageButton = document.createElement('button');
    messageButton.innerText = 'Ok';
    messageButton.onclick = () => messageDialog.close();
    messageDialog.appendChild(messageButton);

    exports.showMessage = (message) => {
        if(!message)//Mensagem vazia é desnecessária, exemplo: redirect
            return;

        messageP.innerText = message;
        messageDialog.showModal();
    };


    const confirmDialog = document.createElement('dialog');
    document.body.appendChild(confirmDialog);
    const confirmP = document.createElement('p');
    confirmDialog.appendChild(confirmP);
    const confirmYesButton = document.createElement('button');
    confirmYesButton.innerText = 'Sim';
    confirmDialog.appendChild(confirmYesButton);
    const confirmNoButton = document.createElement('button');
    confirmNoButton.innerText = 'Não';
    confirmDialog.appendChild(confirmNoButton);

    exports.showConfirm = async (message) => {
        confirmP.innerText = message;
        confirmDialog.showModal();

        return new Promise((res) => {
            confirmYesButton.onclick = () => {
                res(true);
                confirmDialog.close();
            };
            confirmNoButton.onclick = () => {
                res(false);
                confirmDialog.close();
            };
        });
    };

    return exports;
})();
