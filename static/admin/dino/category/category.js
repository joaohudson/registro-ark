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

async function postRegion(name){
    const body = JSON.stringify({name});
    const response = await fetch('/api/dino/category/region', {method: 'POST', body: body});

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

async function fetchCategories(){
    const response = await fetch('/api/dino/categories');

    if(!response.ok){
        throw await response.text();
    }

    return await response.json();
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

    foodList.innerHTML = '';
    for(const food of categories.foods){
        const li = document.createElement('li');
        li.innerText = food.name;
        foodList.appendChild(li);
    }

    regionList.innerHTML = '';
    for(const region of categories.regions){
        const li = document.createElement('li');
        li.innerText = region.name;
        regionList.appendChild(li);
    }

    locomotionList.innerHTML = '';
    for(const locomotion of categories.locomotions){
        const li = document.createElement('li');
        li.innerText = locomotion.name;
        locomotionList.appendChild(li);
    }
}

async function main(){
    await loadData();
}

locomotionButton.onclick = async () => {
    try{
        await postLocomotion(locomotionField.value);
        await loadData();
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
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }

};

main();