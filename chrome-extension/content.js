
$(document).ready(function() {
  const table = $("#DataGrid1");
  const data = [];

  table.find("tr").each(function() {
    const row = [];
    $(this).find("td").each(function() {
      row.push($(this).text().trim());
    });

    const url = $(this).find("td:first-child a").attr("href");
    row.push(url);
    data.push(row);

  });

  data.shift(); // Remove header row

  const updated = data.map(row => {
    const name = row[0].split(', ').reverse().join(' ');

    return {
      name: name,
      gender: row[1] === "M" ? "Male" : "Female",
      city: row[2],
      state: row[3],
      rating: row[4],
      rating_type: row[6] !== "" ? row[6] : "None",
      rating_date: row[5],
      usta_id: decodeURIComponent(row[7].match(/par1=([^&]*)/)?.[1]) || null,
    }
  });

  const urlParams = new URLSearchParams(window.location.search);

  const payload = {
    sectionId: urlParams.get('SectionNodeID'),
    districtId: urlParams.get('DistrictNodeID'),
    areaId: urlParams.get('SubDistrictNodeID'),
    players: updated
  };

  // console.log(payload);
  chrome.runtime.sendMessage({ action: "postData", payload: payload }, (response) => {
    if (response.status === "success") {
      console.log("Data successfully sent to mytennis.team!");
    } else {
      console.error("There was an error sending the data.");
    }
  });
});
