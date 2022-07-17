(function(){

const infoDiv = document.getElementById('itemsDiv');

async function fetchAdms(){
    return await adminFetch('/api/adms', {}, true);
}

async function putAdmPermission(request){
    const body = JSON.stringify(request);
    return await adminFetch('/api/adm/permissions', {method: 'PUT', body});
}

async function main(){
    await loadData();
}

main();


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

        div.appendChild(createLabelDiv(adm));

        div.appendChild(createCheckButtons(adm));

        div.appendChild(createButtons(adm));
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

function createCheckButtons(adm){
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
    };
    checkDiv.appendChild(permDinoCheck);

    const permCategory = document.createElement("label");
    permCategory.innerText = "Gerenciar Categoria";
    checkDiv.appendChild(permCategory);
    const permiCategoryCheck = document.createElement("input");
    permiCategoryCheck.type = "checkbox";
    permiCategoryCheck.checked = adm.permissionManagerCategory;
    permiCategoryCheck.disabled = adm.permissionManagerAdm;
    permiCategoryCheck.onclick = () => {
      adm.permissionManagerCategory = permiCategoryCheck.checked;
    };
    checkDiv.appendChild(permiCategoryCheck);

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

function createButtons(adm){
    const buttonDiv = document.createElement('div');
    buttonDiv.className = 'buttonDiv';
    infoDiv.appendChild(buttonDiv);

    const saveButton = document.createElement('button');
    saveButton.innerText = '💾';
    saveButton.onclick = async () => {
        try{
            await putAdmPermission(adm);
        }
        catch(e){
            dialog.showMessage(e);
            return
        }
        
    }
    buttonDiv.appendChild(saveButton);

    const cancelButton = document.createElement('button');
    cancelButton.innerText = '❌';
    cancelButton.onclick = async () => await loadData();
    buttonDiv.appendChild(cancelButton);

    return buttonDiv;
}

})();