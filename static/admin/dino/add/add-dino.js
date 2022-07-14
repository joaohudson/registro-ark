const regionField = document.getElementById('regionField');
const locomotionField = document.getElementById('locomotionField');
const foodField = document.getElementById('foodField');
const utitlityField = document.getElementById('utilityField');
const trainingField = document.getElementById('trainingField');
const nameField = document.getElementById('nameField');
const createDinoButton = document.getElementById('createDinoButton');

async function fetchCategories(){
    const response = await fetch('/api/dino/categories');

    if(!response.ok){
        throw await response.text();
    }

    return await response.json();
}

async function postDino(dino){
    return adminFetch('/api/dino', {method: 'POST', body: JSON.stringify(dino)});
}

async function createDino(){
    const dino = {
        name: nameField.value,
        utility: utitlityField.value,
        training: trainingField.value,
        regionId: new Number(regionField.value),
        locomotionId: new Number(locomotionField.value),
        foodId: new Number(foodField.value)
    };

    try{
        await postDino(dino);
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }

    dialog.showMessage('Dino criado com sucesso!');
}

function pupulateDropdown(data, dropdown){
    const keys = data.map(k => k.name);
    const values = data.map(v => v.id);
    const optionsKeys =  keys;
    const optionsValues = values;

    for(let i = 0; i < optionsKeys.length; i++){
        const key = optionsKeys[i];
        const value = optionsValues[i];
        const opt = document.createElement('option');
        opt.innerText = key;
        opt.value = value;
        dropdown.appendChild(opt);
    }
}

async function main(){
    let categories;

    try{
        categories = await fetchCategories();
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }

    pupulateDropdown(categories.regions, regionField);
    pupulateDropdown(categories.locomotions, locomotionField);
    pupulateDropdown(categories.foods, foodField);
}

main();


createDinoButton.onclick = createDino;