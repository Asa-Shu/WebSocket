document.addEventListener('DOMContentLoaded', () => {
  let loc = window.location;
  // プロトコルがHTTPかHTTPSかに応じ、'ws:'か'wss:'とする
  let uri = 'ws:';
  if (loc.protocol === 'https:') {
      uri = 'wss:';
  }
  uri += '//' + loc.host + '/ws';
  // console.log(uri)
  // ws://127.0.0.1:8080/ws

  const ws = new WebSocket(uri)
  // WebSocketオブジェクトのonopenイベントに対して、アロー関数を設定
  // onopenイベントは、WebSocketオブジェクトがサーバとの接続を確立したときに発生
  // このイベントが発生する前にWebSocketオブジェクトの他のメソッドを呼び出すと、エラーが発生する場合がある
  ws.onopen = () => {
      console.log('Connected')
  }

  // このアロー関数は、WebSocketオブジェクトからデータを受信したときに実行される
  ws.onmessage = (evt) => {
      let out = document.getElementById('output');
      out.innerHTML += evt.data + '<br>';
  }

  // WebSocketサーバにデータを送信
  const btn = document.querySelector('.btn')
  btn.addEventListener('click', () => {
      ws.send(document.getElementById('input').value)
  })
});
