(function(){

const infoDiv = document.getElementById('itemsDiv');

async function fetchAdms(){
    return await adminFetch('/api/adms', {}, true);
}

async function main(){
    let adms;

    try{
        adms = await fetchAdms();
    }
    catch(e){
        dialog.showMessage(adms);
        return;
    }

    const table = document.createElement('table');
    infoDiv.appendChild(table);

    let tr = document.createElement('tr');
    table.appendChild(tr);
    
    let th = document.createElement('th');
    th.innerText = 'Nome';
    tr.appendChild(th);

    th = document.createElement('th');
    th.innerText = 'Gerenciar Dinos';
    tr.appendChild(th);

    th = document.createElement('th');
    th.innerText = 'Gerenciar Categorias';
    tr.appendChild(th);

    th = document.createElement('th');
    th.innerText = 'Gerenciar Administradores';
    tr.appendChild(th);

    for(const adm of adms){
        tr = document.createElement('tr');

        let td = document.createElement('td');
        td.innerText = adm.name;
        tr.appendChild(td);

        td = document.createElement('td');
        td.innerText = adm.permissionManagerDino ? "Habilitado" : "Desabilitado";
        tr.appendChild(td);

        td = document.createElement('td');
        td.innerText = adm.permissionManagerCategory ? "Habilitado" : "Desabilitado";
        tr.appendChild(td);

        td = document.createElement('td');
        td.innerText = adm.permissionManagerAdm ? "Habilitado" : "Desabilitado";
        tr.appendChild(td);

        table.appendChild(tr);
    }

}

main();

})();