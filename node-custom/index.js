/**
const handle = async (context, body) => {
  return { body: 'first revision' };
}
module.exports = { handle };
*/
function invoke (context,body) {
    return {body: 'first revision' };
}
module.exports = invoke;
