const itemsDiv = document.getElementById('itemsDiv');
const ownerField = document.getElementById('ownerField');
const nameField = document.getElementById('nameField');
const ownerButton = document.getElementById('ownerButton');

async function fetchDinos(owner, name){
    const response = await fetch('/dino?owner=' + owner + '&name=' + name);

    if(!response.ok){
        throw await response.text();
    }

    return await response.json()
}

async function fetchDinoNames(){
    const response = await fetch('/dino/name');

    if(!response.ok){
        throw await response.text();
    }

    return await response.json()
}

async function fetchOwners(){
    const response = await fetch('/owner');

    if(!response.ok){
        throw await response.text();
    }

    return await response.json()
}

async function loadData(){
    let data;
    try{
        data = await fetchDinos(ownerField.value, nameField.value);
    }catch(e){
        alert(e);
        return;
    }

    itemsDiv.innerHTML = '';

    const table = document.createElement('table');

    const fields = [
      { name: "Name", label: "Nome" },
      { name: "Action", label: "Ação" },
      { name: "Owner", label: "Dono" },
    ];

    const trTitle = document.createElement("tr");
    for(const f of fields){
        const th = document.createElement('th');
        th.innerText = f.label;
        trTitle.appendChild(th);
    }
    table.appendChild(trTitle);
    
    for(const obj of data){
        const tr = document.createElement('tr');
        for(const f of fields){
            const td = document.createElement('td');
            td.innerText = obj[f.name];
            tr.appendChild(td);
        }
        table.appendChild(tr);
    }

    itemsDiv.appendChild(table);
}

async function createOwnersDropdown(){
    let owners;
    try{
        owners = await fetchOwners();
    }
    catch(e){
        alert(e);
        return;
    }

    const optionsKeys =  ['Todos', ...owners];
    const optionsValues = ['', ...owners];

    for(let i = 0; i < optionsKeys.length; i++){
        const key = optionsKeys[i];
        const value = optionsValues[i];
        const opt = document.createElement('option');
        opt.innerText = key;
        opt.value = value;
        ownerField.appendChild(opt);
    }
}

async function createDinoNamesDropdown(){
    let names;
    try{
        names = await fetchDinoNames();
    }
    catch(e){
        alert(e);
        return;
    }

    const optionsKeys =  ['Todos', ...names];
    const optionsValues = ['', ...names];

    for(let i = 0; i < optionsKeys.length; i++){
        const key = optionsKeys[i];
        const value = optionsValues[i];
        const opt = document.createElement('option');
        opt.innerText = key;
        opt.value = value;
        nameField.appendChild(opt);
    }
}

async function main(){
    await createOwnersDropdown();
    await createDinoNamesDropdown();
}

main();

ownerButton.onclick = loadData;