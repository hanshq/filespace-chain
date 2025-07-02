# Discovery Questions for CSV Export Feature

These questions will help understand the problem space and user needs for exporting hosting contract data to CSV format.

## Q1: Will this CSV export be used primarily by non-technical users (e.g., for accounting/reporting)?
**Default if unknown:** Yes (CSV exports are typically for business users who need data in spreadsheets)

## Q2: Should the export include related data from hosting inquiries and offers (not just contract IDs)?
**Default if unknown:** Yes (users typically want complete information rather than just ID references)

## Q3: Will users need to filter the exported data (e.g., by date range, creator, or status)?
**Default if unknown:** Yes (filtering is a common requirement for data exports to manage large datasets)

## Q4: Should the CSV export be available through the command-line interface (CLI)?
**Default if unknown:** Yes (the blockchain already has a robust CLI and this fits the existing pattern)

## Q5: Do users need the ability to export all contracts at once without pagination limits?
**Default if unknown:** No (large exports can be problematic; batch/paginated export is safer)