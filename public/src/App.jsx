import React from "react";
import Modal from "./Modal.jsx";
import CanvasContainer from "./CanvasContainer.jsx";

class App extends React.Component {
  constructor() {
    super();
  }
  render() {
    return (
      <div id="root">
        <CanvasContainer />
        <Modal />
      </div>
    );
  }
}
export default App;