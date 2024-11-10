import "./top-bar.styles.scss";
import PageActionButton from "../../components/page-action-buttons/page-action-buttons.component";

const TopbarActions = () => {
  const handleUnsubscribe = () => {
    console.log("Unsubscribed!");
  };

  const handleLabels = () => {
    console.log("Labels");
  };

  const handleBulkDelete = () => {
    console.log("Bulk Deleted");
  };

  return (
    <div className="top-bar-actions-container">
      <PageActionButton type="unsubscribe" handleClick={handleUnsubscribe}>
        Unsubscribe
      </PageActionButton>
      <PageActionButton type="labels" handleClick={handleLabels}>
        Labels
      </PageActionButton>
      <PageActionButton type="bulkDelete" handleClick={handleBulkDelete}>
        Bulk Delete
      </PageActionButton>
    </div>
  );
};

export default TopbarActions;
