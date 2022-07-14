(function(){

const nameField = document.getElementById('nameField');
const passwordField = document.getElementById('passwordField');
const loginButton = document.getElementById('loginButton');

async function login(name, password){
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

loginButton.onclick = async () => {

    const name = nameField.value;
    const password = passwordField.value;
    try{
        const token = await login(name, password);
        localStorage.setItem('token', token);
        location.href = '/admin';
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }
};

})();