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

    const ul = document.createElement("ul");
    for(const obj of data){
        const li = document.createElement("li");
        li.innerText = "Nome: " + obj.Name + " | Ação: " + obj.Action + " | Dono: " + obj.Owner;
        ul.appendChild(li);
    }

    itemsDiv.appendChild(ul);
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