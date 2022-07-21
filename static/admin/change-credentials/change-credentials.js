(function (){

const nameField = document.getElementById('nameField');
const passwordField = document.getElementById('passwordField');
const newNameField = document.getElementById('newNameField');
const newPasswordField = document.getElementById('newPasswordField');
const confirmNewPasswordField = document.getElementById('confirmNewPasswordField');
const saveButton = document.getElementById('saveButton');

async function putAdmCredentials(request){
    const body = JSON.stringify(request);
    return await adminFetch('/api/adm/credentials', {method: 'PUT', body});
}

saveButton.onclick = async () => {
    try{
        if(newPasswordField.value != confirmNewPasswordField.value){
            dialog.showMessage('A senha deve ser igual a confirmada!');
            return
        }
        const request = {
            name: nameField.value,
            password: passwordField.value,
            newName: newNameField.value,
            newPassword: newPasswordField.value
        };
        await putAdmCredentials(request);
        dialog.showMessage('Credenciais alteradas com sucesso!');
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }
}

})();