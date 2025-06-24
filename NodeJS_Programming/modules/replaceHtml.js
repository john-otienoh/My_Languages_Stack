module.exports = function (template, product){ 
    let output = template.replace('{{%IMAGE%}}', product.images);
    output = output.replace('{{%DESCRIPTION%}}', product.description);
    output = output.replace('{{%NAME%}}', product.title);
    output = output.replace('{{%STOCK%}}', product.stock);
    output = output.replace('{{%RATING%}}', product.rating);
    output = output.replace('{{%PRICE%}}', product.price);
    output = output.replace('{{%ID%}}', product.id);
    output = output.replace('{{%DISC%}}', product.discountPercentage)

    return output;

}