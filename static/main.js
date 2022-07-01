const itemsDiv = document.getElementById('itemsDiv');
const regionField = document.getElementById('regionField');
const locomotionField = document.getElementById('locomotionField');
const foodField = document.getElementById('foodField');
const filterButton = document.getElementById('filterButton');

async function fetchDinos(region = '', locomotion = '', food = '', name = ''){
    const response = await fetch('/api/dinos?region=' + region + '&locomotion=' + locomotion + '&food=' + food + '&name=' + name);

    if(!response.ok){
        throw await response.text();
    }

    return await response.json()
}

async function fetchCategories(){
    const response = await fetch('/api/dino/categories');

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
        tr.onclick = () => {
            location.href = '/dino?id=' + obj.id;
        };
        for(const f of categoryNames){
            const td = document.createElement('td');
            td.innerText = obj[f.name];
            tr.appendChild(td);
        }
        table.appendChild(tr);
    }

    itemsDiv.appendChild(table);
}

function pupulateDropdown(data, dropdown){
    const keys = data.map(k => k.name);
    const values = data.map(v => v.id);
    const optionsKeys =  ['Todos', ...keys];
    const optionsValues = ['', ...values];

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
        alert(e);
        return;
    }
    
    pupulateDropdown(categories.regions, regionField);
    pupulateDropdown(categories.locomotions, locomotionField);
    pupulateDropdown(categories.foods, foodField);
}

main();

filterButton.onclick = loadData;