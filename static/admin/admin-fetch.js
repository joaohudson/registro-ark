const adminFetch = async (url, options, responseJson) => {
    const token = localStorage.getItem('token');
    if(!options.headers){
        options.headers = new Headers();
    }
    options.headers.append('Authorization', token);
    const response = await fetch(url, options);

    if(!response.ok){
        if(response.status == 401){
            location.href = '/admin/login';
            throw await response.text();
        }
        else if(response.status == 403){
            location.href = '/admin';
            throw await response.text();
        }
        else{
            throw await response.text();
        }
    }

    return responseJson ? await response.json() : await response.text();
};