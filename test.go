const startDateQueryResult = [["2023-02-1 00:00:00.000 +00:00"]];
const startDate = startDateQueryResult[0][0];

let endDate = null;

if (startDate) {
  const sDate = new Date(startDate);
  const extractedDate = sDate.getDate();

  const startMonth = sDate.getMonth();
  const startYear = sDate.getFullYear();
  let endMonth = startMonth;
  let endYear = startYear;

  // Calculate end date
  const nextMonth = new Date(startYear, startMonth + 1, 1);
  nextMonth.setHours(nextMonth.getHours() - 1); // Get the last moment of the current month
  const maxDaysInStartMonth = nextMonth.getDate();

  // Handle months with 30 days
  if ([3, 5, 8, 10].includes(startMonth)) {
    endMonth = (startMonth + 1) % 12;
    if (endMonth === 0) {
      endYear++;
    }
    endDate = new Date(endYear, endMonth, extractedDate);
  }
  // Handle months with 31 days and February
  else {
    let endDateDay;
    if (startMonth === 1 && isLeapYear(startYear)) {
      endDateDay = Math.min(1, 29, maxDaysInStartMonth, 30);
      endMonth = (startMonth + 1) % 12;
    } else {
      endDateDay = Math.min(1, 28, maxDaysInStartMonth, 30);
      endMonth = (startMonth + 1) % 12;
    }
    if (endMonth === 0) {
      endYear++;
    }
    endDate = new Date(endYear, endMonth, endDateDay);
  }
}

console.log(endDate.toString());

function isLeapYear(year: number): boolean {
  return (year % 4 === 0 && year % 100 !== 0) || year % 400 === 0;
}
