const locomotionList = document.getElementById('locomotionList');
const regionList = document.getElementById('regionList');
const foodList = document.getElementById('foodList');

const locomotionField = document.getElementById('locomotionField');
const regionField = document.getElementById('regionField');
const foodField = document.getElementById('foodField');

const locomotionButton = document.getElementById('locomotionButton');
const regionButton = document.getElementById('regionButton');
const foodButton = document.getElementById('foodButton');

async function postLocomotion(name){
    const body = JSON.stringify({name});
    const response = await fetch('/api/dino/category/locomotion', {method: 'POST', body: body});

    if(!response.ok){
        throw await response.text();
    }
}

async function deleteLocomotion(id){
    const response = await fetch('/api/dino/category/locomotion?id=' + id, {method: 'DELETE'});

    if(!response.ok){
        throw await response.text();
    }
}

async function postRegion(name){
    const body = JSON.stringify({name});
    const response = await fetch('/api/dino/category/region', {method: 'POST', body: body});

    if(!response.ok){
        throw await response.text();
    }
}

async function deleteRegion(id){
    const response = await fetch('/api/dino/category/region?id=' + id, {method: 'DELETE'});

    if(!response.ok){
        throw await response.text();
    }
}

async function postFood(name){
    const body = JSON.stringify({name});
    const response = await fetch('/api/dino/category/food', {method: 'POST', body: body});

    if(!response.ok){
        throw await response.text();
    }
}

async function deleteFood(id){
    const response = await fetch('/api/dino/category/food?id=' + id, {method: 'DELETE'});

    if(!response.ok){
        throw await response.text();
    }
}

async function fetchCategories(){
    const response = await fetch('/api/dino/categories');

    if(!response.ok){
        throw await response.text();
    }

    return await response.json();
}

function populateListCategories(data, list, deleteFunc){
    list.innerHTML = '';
    for(const category of data){
        const li = document.createElement('li');
        const span = document.createElement('span');
        span.innerText = category.name;
        li.appendChild(span);
        const button = document.createElement('button');
        button.innerText = '❌';
        button.onclick = async () => {
            try{
                const ok = await dialog.showConfirm('Tem certeza que deseja deletar ' + category.name + ' ?');
                if(ok)
                    await deleteFunc(category.id);
                await loadData();
            }
            catch(e){
                dialog.showMessage(e);
                return;
            }
        };
        li.appendChild(button);
        list.appendChild(li);
    }
}

async function loadData(){
    let categories;

    try{
        categories = await fetchCategories();
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }

    populateListCategories(categories.foods, foodList, deleteFood);
    populateListCategories(categories.regions, regionList, deleteRegion);
    populateListCategories(categories.locomotions, locomotionList, deleteLocomotion);
}

async function main(){
    await loadData();
}

locomotionButton.onclick = async () => {
    try{
        await postLocomotion(locomotionField.value);
        await loadData();
        locomotionField.value = '';
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }
};

regionButton.onclick = async () => {
    try{
        await postRegion(regionField.value);
        await loadData();
        regionField.value = '';
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }
};

foodButton.onclick = async () => {
    try{
        await postFood(foodField.value);
        await loadData();
        foodField.value = '';
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }

};

main();
