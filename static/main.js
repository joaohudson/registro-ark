const itemsDiv = document.getElementById('itemsDiv');
const regionField = document.getElementById('regionField');
const locomotionField = document.getElementById('locomotionField');
const foodField = document.getElementById('foodField');
const filterButton = document.getElementById('filterButton');

async function fetchDinos(region = '', locomotion = '', food = '', name = ''){
    const response = await fetch('/api/dino?region=' + region + '&locomotion=' + locomotion + '&food=' + food + '&name=' + name);

    if(!response.ok){
        throw await response.text();
    }

    return await response.json()
}

async function fetchCategories(){
    const response = await fetch('/api/dino/category');

    if(!response.ok){
        throw await response.text();
    }

    return await response.json();
}

async function loadData(){
    let data;
    try{
        data = await fetchDinos(regionField.value, locomotionField.value, foodField.value);
    }catch(e){
        alert(e);
        return;
    }

    itemsDiv.innerHTML = '';

    const table = document.createElement('table');

    const categoryNames = [
      { name: "name", label: "Nome" },
      { name: "food", label: "Alimentação" },
      { name: "locomotion", label: "Locomoção" },
      { name: "region", label: "Região" }
    ];

    const trTitle = document.createElement("tr");
    for(const f of categoryNames){
        const th = document.createElement('th');
        th.innerText = f.label;
        trTitle.appendChild(th);
    }
    table.appendChild(trTitle);
    
    for(const obj of data){
        const tr = document.createElement('tr');
        for(const f of categoryNames){
            const td = document.createElement('td');
            td.innerText = obj[f.name];
            tr.appendChild(td);
        }
        table.appendChild(tr);
    }

    itemsDiv.appendChild(table);
}

function createFoodDropdown(categories){
    const {foods} = categories;
    const optionsKeys =  ['Todos', ...foods];
    const optionsValues = ['', ...foods];

    for(let i = 0; i < optionsKeys.length; i++){
        const key = optionsKeys[i];
        const value = optionsValues[i];
        const opt = document.createElement('option');
        opt.innerText = key;
        opt.value = value;
        foodField.appendChild(opt);
    }
}

function createLocomotionDropdown(categories){
    const {locomotions} = categories;
    const optionsKeys =  ['Todos', ...locomotions];
    const optionsValues = ['', ...locomotions];

    for(let i = 0; i < optionsKeys.length; i++){
        const key = optionsKeys[i];
        const value = optionsValues[i];
        const opt = document.createElement('option');
        opt.innerText = key;
        opt.value = value;
        locomotionField.appendChild(opt);
    }
}

function createRegionDropdown(categories){
    const {regions} = categories;
    const optionsKeys =  ['Todos', ...regions];
    const optionsValues = ['', ...regions];

    for(let i = 0; i < optionsKeys.length; i++){
        const key = optionsKeys[i];
        const value = optionsValues[i];
        const opt = document.createElement('option');
        opt.innerText = key;
        opt.value = value;
        regionField.appendChild(opt);
    }
}

async function main(){
    
    let categories;

    try{
        categories = await fetchCategories();
    }
    catch(e){
        alert(e);
        return;
    }
    
    createRegionDropdown(categories);
    createLocomotionDropdown(categories);
    createFoodDropdown(categories);
}

main();

filterButton.onclick = loadData;