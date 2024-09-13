import Header from "../header/header.component";
import CustomButton from "../button/button.component";
import "./popup.styles.scss";

const Popup = () => (
  <div className="popup border-2 border-sky-500 px-4 py-4 w-full">
    <Header />

    {/* Overview section */}
    <div className="overview py-4 px-4">
      <div className="header w-full py-2">
        <h2 className="title font-semibold">Inbox Overview</h2>
        <CustomButton>Icon Settings</CustomButton>
      </div>

      <div className="overview-content py-2">
        <p className="unread mr-12">Unread: 15</p>
        <p className="important">Imporant: 4</p>
      </div>
    </div>
    <hr className="h-px bg-slate-300" />
  </div>
);

export default Popup;
