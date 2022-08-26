(function(){

const infoDiv = document.getElementById('itemsDiv');

//registro
const userField = document.getElementById('userField');
const passwordField = document.getElementById('passwordField');
const confirmPasswordField = document.getElementById('confirmPasswordField');
const singupButton = document.getElementById('singupButton');
const managerDinoCheck = document.getElementById('managerDinoCheck');
const managerCategoryCheck = document.getElementById('managerCategoryCheck');

async function fetchAdms(){
    return await adminFetch('/api/adms', {}, true);
}

async function putAdmPermission(request){
    const body = JSON.stringify(request);
    return await adminFetch('/api/adm/permissions', {method: 'PUT', body});
}

async function postAdm(request){
    const body = JSON.stringify(request);
    return await adminFetch('/api/adm', {method: 'POST', body});
}

async function deleteAdm(id){
    return await adminFetch('/api/adm?id='+id, {method: 'DELETE'});
}

async function main(){
    await loadData();
}

main();

singupButton.onclick = async () => {
    try{
        if(passwordField.value != confirmPasswordField.value){
            dialog.showMessage('A senha deve ser igual a confirmada!');
            return;
        }

        const admRequest = {
            name: userField.value,
            password: passwordField.value,
            permissionManagerDino: managerDinoCheck.checked,
            permissionManagerCategory: managerCategoryCheck.checked
        };

        await postAdm(admRequest);
        userField.value = '';
        passwordField.value = '';
        confirmPasswordField.value = '';
        managerDinoCheck.checked = false;
        managerCategoryCheck.checked = false;
        await loadData();
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }
};


async function loadData(){
    let adms;

    try{
        adms = await fetchAdms();
    }
    catch(e){
        dialog.showMessage(adms);
        return;
    }

    infoDiv.innerHTML = '';

    for(const adm of adms){
        const div = document.createElement('div');
        div.className = 'itemDiv';
        infoDiv.appendChild(div);

        const channel = {
            onmodified: () => {}
        };

        div.appendChild(createLabelDiv(adm));

        div.appendChild(createCheckButtons(adm, channel));

        div.appendChild(createButtons(adm, channel));

        div.oncontextmenu = async (e) => {
            e.preventDefault();
            const ok = await dialog.showConfirm('Deseja remover ' + adm.name + ' ?');
            if(ok){
                try{
                    await deleteAdm(adm.id);
                    await loadData();
                }
                catch(e){
                    dialog.showMessage(e);
                    return;
                }
            }
        }
    }
}

function createLabelDiv(adm){
    const labelDiv = document.createElement("div");
    labelDiv.className = "centered";
    const nameLabel = document.createElement("label");
    nameLabel.innerText = adm.name;
    labelDiv.appendChild(nameLabel);

    return labelDiv;
}

function createCheckButtons(adm, channel){
    const checkDiv = document.createElement("div");
    checkDiv.className = "centered";

    const permDinoLabel = document.createElement("label");
    permDinoLabel.innerText = "Gerenciar Dino";
    checkDiv.appendChild(permDinoLabel);
    const permDinoCheck = document.createElement("input");
    permDinoCheck.type = "checkbox";
    permDinoCheck.checked = adm.permissionManagerDino;
    permDinoCheck.disabled = adm.permissionManagerAdm;
    permDinoCheck.onclick = () => {
      adm.permissionManagerDino = permDinoCheck.checked;
      channel.onmodified();
    };
    checkDiv.appendChild(permDinoCheck);

    const permCategory = document.createElement("label");
    permCategory.innerText = "Gerenciar Categoria";
    checkDiv.appendChild(permCategory);
    const permCategoryCheck = document.createElement("input");
    permCategoryCheck.type = "checkbox";
    permCategoryCheck.checked = adm.permissionManagerCategory;
    permCategoryCheck.disabled = adm.permissionManagerAdm;
    permCategoryCheck.onclick = () => {
      adm.permissionManagerCategory = permCategoryCheck.checked;
      channel.onmodified();
    };
    checkDiv.appendChild(permCategoryCheck);

    const permAdmLabel = document.createElement("label");
    permAdmLabel.innerText = "Gerenciar Administrador";
    checkDiv.appendChild(permAdmLabel);
    const permAdmCheck = document.createElement("input");
    permAdmCheck.type = "checkbox";
    permAdmCheck.checked = adm.permissionManagerAdm;
    permAdmCheck.disabled = true;
    checkDiv.appendChild(permAdmCheck);

    return checkDiv;
}

function createButtons(adm, channel){
    const buttonDiv = document.createElement('div');
    buttonDiv.className = 'buttonDiv';
    infoDiv.appendChild(buttonDiv);

    const saveButton = document.createElement('button');
    saveButton.innerText = '💾';
    saveButton.title = 'Salvar alterações';
    saveButton.disabled = true;
    buttonDiv.appendChild(saveButton);
    const cancelButton = document.createElement('button');
    cancelButton.disabled = true;
    cancelButton.innerText = '↩️';
    cancelButton.title = 'Desfazer alterações';
    buttonDiv.appendChild(cancelButton);


    saveButton.onclick = async () => {
        try{
            await putAdmPermission(adm);
            saveButton.disabled = true;
            cancelButton.disabled = true;
        }
        catch(e){
            dialog.showMessage(e);
            return
        }
        
    }

    cancelButton.onclick = async () => {
        await loadData();
        saveButton.disabled = true;
        cancelButton.disabled = true;
    }

    channel.onmodified = () => {
        saveButton.disabled = false;
        cancelButton.disabled = false;
    };

    return buttonDiv;
}

})();