const itemsDiv = document.getElementById('itemsDiv');
const ownerField = document.getElementById('ownerField');
const ownerButton = document.getElementById('ownerButton');

async function fetchDinos(owner){
    const response = await fetch('/dino?owner=' + owner);

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
        data = await fetchDinos(ownerField.value);
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

async function main(){
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

main();

ownerButton.onclick = loadData;