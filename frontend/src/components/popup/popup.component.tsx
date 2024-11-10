import Header from "../header/header.component";
import CustomButton from "../button/button.component";
import "./popup.styles.scss";

const Popup = () => (
  <div className="popup">
    <Header />

    {/* Overview section */}
    <div className="overview py-8 px-8">
      <div className="header w-full py-4">
        <h2 className="title font-semibold">Inbox Overview</h2>
        <CustomButton>Icon Settings</CustomButton>
      </div>

      <div className="overview-content py-2">
        <p className="unread mr-12">Unread: 15</p>
        <p className="important">Imporant: 4</p>
      </div>
    </div>
    <hr className="h-px bg-slate-300" />

    {/* Category overview section */}
    <div className="category-overview px-8 py-4">
      <h2 className="title font-semibold">Quick Actions:</h2>
      <div className="action-buttons pt-4">
        <CustomButton type="work">Work (5)</CustomButton>
        <CustomButton type="personal">Personal (6)</CustomButton>
        <CustomButton type="promotional">Promotional (3)</CustomButton>
        <CustomButton type="spam">Spam (13)</CustomButton>
        <CustomButton type="social">Social (5)</CustomButton>
        <CustomButton type="all-mails">All Mails</CustomButton>
      </div>
    </div>
    <hr className="h-px bg-slate-300" />

    {/* Important emails section */}
    <div className="important-imails px-8 py-4">
      <h2 className="title font-semibold">Important Emails:</h2>
      <ol className="list-decimal pl-8 pt-2">
        <li>Attend a zoom session with John Harris</li>
        <li>Unsubcribe HULU before deadline</li>
        <li>Mymathlab deadline is tomorrow.</li>
      </ol>
    </div>
    <hr className="h-px bg-slate-300" />

    {/* Quick actions section */}
    <div className="quick-actions px-8 py-4 pb-8">
      <h2 className="title font-semibold">Quick Actions:</h2>
      <div className="actions pt-4">
        <CustomButton>Mark All Read</CustomButton>
        <CustomButton>Delete Spam</CustomButton>
      </div>
    </div>
  </div>
);

export default Popup;
