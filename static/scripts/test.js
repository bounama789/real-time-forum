let form = document.getElementById("form")

form.addEventListener("submit", async (e) => {
    e.preventDefault()
    let f = e.currentTarget
    let url = f.action

    let formFields = new FormData(f)
    idstr = formFields.get("category")
    var data = new Map()
    for ( [key,value] of formFields){
        if (key == "category") {
            key="category_id"
            value = parseInt(value)
        }
        data.set(key,value)
    }
    let res = await postFormFieldsAsJson({ url, data })
})


async function postFormFieldsAsJson({ url, data }) {
    //Create an object from the form data entries
    let formDataObject = Object.fromEntries(data.entries());
    // Format the plain form data as JSON
    let formDataJsonString = JSON.stringify(formDataObject);

    //Set the fetch options (headers, body)
    let fetchOptions = {
        //HTTP method set to POST.
        method: "POST",
        //Set the headers that specify you're sending a JSON body request and accepting JSON response
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        // POST request body as JSON string.
        body: formDataJsonString,
    };

    //Get the response body as JSON.
    //If the response was not OK, throw an error.
    let res = await fetch(url, fetchOptions);

    //If the response is not ok throw an error (for debugging)
    // if (!res.ok) {
    //     let error = await res.text();
    //     throw new Error(error);
    // }
    //If the response was OK, return the response body.
    return res;
}
