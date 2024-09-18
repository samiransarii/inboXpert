import "./header.styles.scss";

const Header = () => {
  const handleClick = () => {
    window.close();
  };

  return (
    <div className="popup-header bg-clr-gray px-8 py-4 pr-14">
      <div className="title font-semibold">InboXpert</div>
      <button onClick={handleClick} className="control-btn-container">
        x
      </button>
    </div>
  );
};

export default Header;
