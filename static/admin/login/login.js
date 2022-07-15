(function(){

const nameField = document.getElementById('nameField');
const passwordField = document.getElementById('passwordField');
const loginButton = document.getElementById('loginButton');

async function postLogin(name, password){
    const body = JSON.stringify({
        name,
        password
    });
    const response = await fetch('/api/adm/login', {method: 'POST', body});

    if(!response.ok){
        throw await response.text();
    }

    return await response.text();
}

async function login(){
    const name = nameField.value;
    const password = passwordField.value;
    try{
        const token = await postLogin(name, password);
        localStorage.setItem('token', token);
        location.href = '/admin';
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }
}

loginButton.onclick = login;

passwordField.onkeydown = async (e) => {
    if(e.key == 'Enter'){
        await login();
    }
};

})();