
var DeviceLayout = function(data) {
	
    this.init = function() {
    	if ( this.cells && this.cells.content ) {
	    	for ( var index=0; index<this.cells.content.length; index++) {
	    		var cell = this.cells.content[index];
	    		var device_cell = new DeviceCell(cell);
	    		this.cells.content[index] = device_cell;
	    	}
    	} else {
    		if ( this.cells ) {
    			this.cells.content = [];
    		} else {
    			this.cells = {};
    			this.cells.content = [];
    		}
    		
    	}
    	this.setCurrentPageId("1");
    };
    
    this.reset = function() {
    	
    	var theme_id = this.settings.theme_id;
    	var theme_version = this.settings.theme_version;
    	
    	this.cells.content = [];
    	this.contacts.content = [];
    	this.apps.content = [];
//    	this.settings = {};
    	
    	this.settings.theme_id = theme_id;
    	this.settings.theme_version = theme_version;
    };
	
	this.save = function() {
		Log.d("DeviceLayout.save");
		Server.layout_put(this.device_id, this, {
			onSuccess : function(data) {
				Log.d("App.Server.saveDeviceLayout(): success");
			},
			onFailure : function(data) {
				Log.d("App.Server.saveDeviceLayout(): failure");

			}
		});
	};
	
	this.getCurrentPageId = function() {
		return this.settings.currentPageId;
	};
	
	this.setCurrentPageId = function(pageId) {
		this.settings.currentPageId = pageId;
	};
	
	this.getSelectedCell = function() {
		for ( var index=0; index<this.cells.content.length; index++) {
			var cell = this.cells.content[index];
			if ( cell.isSelected() ) {
				return cell;
			}
		}
		return null;
	};
	
	this.getSiblingCell = function(source_cell) {
		for ( var index=0; index<this.cells.content.length; index++) {
			var cell = this.cells.content[index];
			if ( cell.row == source_cell.row && cell != source_cell ) {
				return cell;
			}
		}
		return null;
		
	};

	this.updateSetting = function(key, value) {
		Log.d("DeviceLayout.updateSetting(" +key + ", " + value);
		this.settings[key] = value;
		
		Server.device_update_setting(this.device_id, key, value, {
			onSuccess : {},
			onFailure : {}
		});

	};
	
	this.getSettingValueOrDefault = function(key, defaultValue) {
		var value = this.settings[key];
		if ( value == null ) {
			return defaultValue;
		} else {
			return value;
		}
	};
		
//	this.toggleApp = function(packageName) {
//		Log.d("DeviceLayout.toggleApp(" + packageName + ")");
//		if (!this.apps ) {
//			this.apps = [];
//		}
//		var page_id = this.getCurrentPageId(); 
//
//		if ( this.apps.content.indexOf(packageName) > -1 ) {
//			// then remove it
//			var index = this.apps.content.indexOf(packageName);
//			this.apps.content.splice(index, 1);
//			
//			var cell_index = this.indexOf({type:"app", "id": app_id, "page_id":page_id});
//			if( cell_index > -1 ) {
//				this.cells.content.splice(cell_index, 1);
//			}
//			Server.device_remove_app(this.device_id, packageName);
//			
//		} else {
//			this.apps.content.push(packageName);
//			var emptyCell = this.getFirstEmptyCell(page_id);
//			var app_id = packageName;
//			Server.device_add_app(this.device_id, packageName);
//			this.moveApp(page_id, emptyCell.row, app_id, emptyCell.position);
//		}
//	};
//
//	this.toggleContact = function(contact_id) {
//		Log.d("DeviceLayout.toggleContact(" + contact_id + ")");
//		if (!this.contacts ) {
//			this.contacts = [];
//		}
//		var page_id = this.getCurrentPageId(); 
//		if ( this.contacts.content.indexOf(contact_id) > -1 ) {
//			var index = this.contacts.content.indexOf(contact_id);
//			this.contacts.content.splice(index, 1);
//			var cell_index = this.indexOf({type:"contact", "contact_id": contact_id, "page_id": page_id });
//			if( cell_index > -1 ) {
//				this.cells.content.splice(cell_index, 1);
//			}
//			Server.device_remove_contact(this.device_id, contact_id, page_id);
//			
//		} else {
//			this.contacts.content.push(contact_id);
//			var emptyCell = this.getFirstEmptyCell(page_id);
//			Server.device_add_contact(this.device_id, contact_id);
//			this.moveContact(page_id, emptyCell.row, contact_id, emptyCell.position);
//		}
//	};
	
	this.moveApp = function(page_id, row, app_id, position) {
		Server.device_move_app(this.device_id, page_id, app_id, row, position);
		var cell = this.getCellByAppId(app_id, page_id);
		if (cell == null) {
			cell = { "type": "app", "page_id": page_id, "row": row, "app_id": app_id, "position": position };
			cell = new DeviceCell(cell);
			this.cells.content.push(cell);
		} else {
			cell.row = row;
			cell.position = position;
		}
	};
	
	this.moveContact = function(page_id, row, contact_id, position) {
		Server.device_move_contact(this.device_id, page_id, contact_id, row, position);
		var cell = this.getCellByContactId(contact_id, page_id);
		if (cell == null) {
			cell = { "type": "contact", "page_id": page_id, "row": row, "contact_id": contact_id, "position": position };
			cell = new DeviceCell(cell);
			this.cells.content.push(cell);
		} else {
			cell.row = row;
			cell.position = position;
		}
	};
	
	/**
	 * returns a cell, or undefined, from an object like { "row": 1, "position": left }
	 */
	this.getCellAt = function(obj) {
		for (var index=0; index<this.cells.content.length; index++) {
			var cell = this.cells.content[index];
			if ( cell.row == obj.row && cell.position == obj.position && cell.page_id == obj.page_id ) {
				return cell;
			}
		}
		return null;
	};
	
	this.getCellById = function(id, page_id) {
		for(var index=0;index<this.cells.content.length;index++) {
			var cell = this.cells.content[index];
			if ( cell.page_id == page_id && ( cell.contact_id == id || cell.app_id == id ) ) {
				return cell;
			}
		}
		return null;
	};
	
	this.getCellByAppId = function(id, page_id) {
		for(var index=0;index<this.cells.content.length;index++) {
			var cell = this.cells.content[index];
			if ( cell.page_id == page_id && cell.type == "app" && cell.app_id == id ) {
				return cell;
			}
		}
		return null;
	};
	
	this.getCellByContactId = function(id, page_id) {
		for(var index=0;index<this.cells.content.length;index++) {
			var cell = this.cells.content[index];
			if ( cell.page_id == page_id && cell.type == "contact" && cell.contact_id == id ) {
				return cell;
			}
		}
		return null;
	};
	
	this.getCellsForRow = function(row, page_id) {
		var cells = [];
		for(var index=0;index<this.cells.content.length;index++) {
			var cell = this.cells.content[index];
			if ( cell.row == row && cell.page_id == page_id) {
				cells.push(cell);
			}
			
		}
		return cells;
		
	};
	
	/**
	 * does what it says, returns an object { "row": 1, "col": "left|right" } indicating the first available
	 * row/column 
	 */
	this.getFirstEmptyCell = function(page_id) {
		
		// build a map to work out whats not taken
		var map = {};
		var max_rows = 10;
		for ( var row = 0; row<max_rows; row++) {
			var key = "" + row + "_left";
			map[key] = false;
			key = "" + row + "_right";
			map[key] = false;
			key = "" + row + "_banner";
			map[key] = false;
		}
		
		// update the map with the taken positions 

		for(var index=0;index<this.cells.content.length;index++) {
			var cell = this.cells.content[index];
			if ( cell.page_id == page_id ) {
				var key = cell.row + "_" + cell.position.toLowerCase();
				map[key] = true;
			}
		}

		for (var row=1; row<max_rows; row++) {
			var key_left = row + "_left";
			var key_right = row + "_right";
			var key_banner = row + "_banner";
			var left_taken = map[key_left];
			var right_taken = map[key_right];
			var banner_taken = map[key_banner];
			
			if ( banner_taken ) { 
				continue;
			} else if ( left_taken && right_taken ) {
				continue;
			} else if ( left_taken ) {
				return { "row": row, "position": "RIGHT" };
			} else if ( right_taken ) {
				return { "row": row, "position": "LEFT" };
			} else {
				// nothing is taken on this row
				return { "row": row, "position": "LEFT" };
			}
			
		}
		
	};
	
	/**
	 * indicates if the contact_id displayed on-page
	 */
	this.hasContact = function(contact_id, page_id) {
		return this.indexOf({type:"contact", "id": contact_id, "page_id": page_id}) > -1
	};
	
	/**
	 * indicates if the contact_id displayed on-page
	 */
	this.hasApp = function(app_id, page_id) {
		return this.indexOf({type:"app", "id": app_id, "page_id": page_id}) > -1
	};
	
	this.addContact = function(contact_id, page_id) {
		var cell = this.getFirstEmptyCell(page_id);
		if ( this.contacts.content.indexOf(contact_id) == -1 ) {
			this.contacts.content.push(contact_id);
			Server.device_add_contact(this.device_id, contact_id);
		}
		this.moveContact(page_id, cell.row, contact_id, cell.position);
	};

	this.addApp = function(app_id, page_id) {
		var cell = this.getFirstEmptyCell(page_id);
		if ( this.apps.content.indexOf(app_id) == -1 ) {
			this.apps.content.push(app_id);
			Server.device_add_app(this.device_id, app_id);
		}
		this.moveApp(page_id, cell.row, app_id, cell.position);
	};

	/**
	 * position of this type in the 
	 */
	this.indexOf = function(obj) {
		for ( var index=0; index<this.cells.content.length; index++) {
			var cell = this.cells.content[index];
			if ( cell.type == obj.type && cell.getId() == obj.id && cell.page_id == obj.page_id) { 
				return index;
			}
		}
		return -1;
	};
	
	this.moveSelectedCellUp = function() {
		Log.d("DeviceLayout.moveSelectedCellUp()");
		var selectedCell = this.getSelectedCell();
		if ( selectedCell == null ) {
			return;
		}
		var page_id = this.getCurrentPageId();
		var device_id = this.device_id;
		var targetRow = new Number(selectedCell.row) - 1;
		var originalRow = new Number(selectedCell.row);
		var targetPosition = selectedCell.position;
		var sibling = this.getSiblingCell(selectedCell);
		var cellsAbove = this.getCellsForRow(targetRow, page_id);
		var cells = [];
		
		if ( selectedCell.isBanner() ) {
			
			if ( cellsAbove.length == 0 ) {
				// nothing above; we cannot move
				return;
			} else {
				// we can swap the rows out
				selectedCell.row = targetRow;
				cells.push(selectedCell);
				for ( var index=0; index<cellsAbove.length; index++) {
					cellsAbove[index].row = originalRow;
					cells.push(cellsAbove[index]);
				}
			}
			
		} else {
		
			if ( cellsAbove.length == 0 ) {
				// nothing above.  It's possible that if we have a sibling, we can reorder everything
				if ( sibling != null ) {
					for ( var index=0; index<this.cells.content.length; index++) {
						if ( this.cells.content[index] == selectedCell ) {
							continue;
						}
						this.cells.content[index].row = new Number(this.cells.content[index].row) + 1;
						cells.push(this.cells.content[index]);
					}
	//				selectedCell.row = new Number(selectedCell.row) - 1;
	//				cells.push(selectedCell);
				} else {
					return;
				}		
				
			} else {
				// there is something above
				// however, lets check to see, is there space directly above?
				var cellDirectlyAbove = this.getCellAt({"row": targetRow, "position": targetPosition, "page_id": page_id});
				if ( cellDirectlyAbove == null ) {
					var bannerAbove = this.getCellAt({"row": targetRow, "position": "BANNER", "page_id": page_id});
					if ( bannerAbove == null ) {
						// excellent; there is nothing directly above us, so we can just move
						selectedCell.row = targetRow;
						cells.push(selectedCell);
					} else {
						// the cell above is a banner; 
						// if we have a sibling, then we want to move up above the sibling
						// if we dont, we can flip the banner and this cell
						if ( sibling == null ) {
							bannerAbove.row = originalRow;
							selectedCell.row = targetRow;
							cells.push(bannerAbove);
							cells.push(selectedCell);
						} else {
							// we have a sibling and above us is the banner.  We need to move
							// the sibling and everthing below us DOWN one
							var newRow = new Number(originalRow) + 1;
							for ( var index=0; index<this.cells.content.length; index++) {
								var candidate = this.cells.content[index];
								if (candidate.page_id == page_id && candidate.row >= originalRow && candidate != selectedCell ) {
									candidate.row = newRow;
									cells.push(candidate);
								}
							} 
							
						}

						
					}
					
				} else {
					// ok, if something is directly above us, just flip its row to be ours and vice versa
					cellDirectlyAbove.row = new Number(cellDirectlyAbove.row) + 1;
					selectedCell.row = targetRow;
					cells.push(cellDirectlyAbove);
					cells.push(selectedCell);
					
				}
				
			}
			
		}

		if ( cells.length > 0 ) {
			this.fixAnyEmptyRows(cells);
			Server.device_move_cells(device_id, cells);
			App.View.showDragDrop();
		}

		
	};

	this.moveSelectedCellDown = function() {
		Log.d("DeviceLayout.moveSelectedCellDown()");
		var selectedCell = this.getSelectedCell();
		if ( selectedCell == null ) {
			return;
		}
		var page_id = this.getCurrentPageId();
		var device_id = this.device_id;
		var originalRow = new Number(selectedCell.row);
		var targetRow = new Number(selectedCell.row) + 1;
		var targetPosition = selectedCell.position;
		var sibling = this.getSiblingCell(selectedCell);
		var cellsBelow = this.getCellsForRow(targetRow, page_id);
		var cells = [];
		
		if ( selectedCell.isBanner() ) {

			if ( cellsBelow.length == 0 ) {
				// nothing above; we cannot move
				return;
			} else {
				// we can swap the rows out
				selectedCell.row = targetRow;
				cells.push(selectedCell);
				for ( var index=0; index<cellsBelow.length; index++) {
					cellsBelow[index].row = originalRow;
					cells.push(cellsBelow[index]);
				}
			}

			
		} else {
			
			if ( cellsBelow.length == 0 ) {
				// nothing below.  We can't go down if we are the only cell in this row as we must be
				// at the end
				if ( sibling == null ) {
					return;
				} else {
					// otherwise we can go down
					selectedCell.row = targetRow;
					cells.push(selectedCell);
				}
			} else {
				// there is something below us.  What do we need to do?
				// check, is there space directly below where we are?
				var cellDirectlyBelow = this.getCellAt({ "row": targetRow, "position": targetPosition, "page_id": page_id});
				if ( cellDirectlyBelow == null ) {
					
					var bannerBelow = this.getCellAt({ "row": targetRow, "position": "BANNER", "page_id": page_id});
					if ( bannerBelow == null ) {
						return;
					} else {
						// there is a banner below us; do we have a sibling
						if ( sibling == null ) {
							// no, just swap the banner below and our position
							// this is where we can flip the two around
							selectedCell.row = targetRow;
							bannerBelow.row = originalRow;
							cells.push(selectedCell);
							cells.push(bannerBelow);
						} else {
							// we are leaving our sibling, what to do is push everthing below us
							// down one row, including ourselves
							var newRow = targetRow; 
							for ( var index=0; index<this.cells.content.length; index++) {
								var candidate = this.cells.content[index];
								if ( candidate.page_id == page_id && candidate.row >= originalRow && candidate != sibling ) {
									candidate.row = new Number(candidate.row) + 1;
									cells.push(candidate);
								}
								
							}
						}
					}
					
//					// great, just move this down
//					selectedCell.row = targetRow;
//					cells.push(selectedCell);
				} else {
					// this is where we can flip the two around
					selectedCell.row = targetRow;
					cellDirectlyBelow.row = new Number(cellDirectlyBelow.row) - 1;
					cells.push(selectedCell);
					cells.push(cellDirectlyBelow);
				}
			}
			
		}
		
		if ( cells.length > 0 ) {
			this.fixAnyEmptyRows(cells);
			Server.device_move_cells(device_id, cells);
			App.View.showDragDrop();
		}
		
	};

	this.moveSelectedCellLeft = function() {
		Log.d("DeviceLayout.moveSelectedCellLeft()");
		var selectedCell = this.getSelectedCell();
		if ( selectedCell == null ) {
			return;
		}
		if (selectedCell.isLeft() ) {
			return;
		}
		
		var sibling = this.getSiblingCell(selectedCell);
		var cells = [];
		var device_id = this.device_id;
		if ( sibling != null ) {
			sibling.position = "RIGHT";
			selectedCell.position = "LEFT";
			cells.push(sibling);
			cells.push(selectedCell);
		} else if ( selectedCell.isRight() ) {
			selectedCell.position = "BANNER";
			cells.push(selectedCell);
		} else if ( selectedCell.isBanner() ) {
			cells.push(selectedCell);
			selectedCell.position = "LEFT";
		}
		
		if ( cells.length > 0 ) {
			Server.device_move_cells(device_id, cells);
			App.View.showDragDrop();
		}
		
	};

	this.moveSelectedCellRight = function() {
		Log.d("DeviceLayout.moveSelectedCellRight()");
		var selectedCell = this.getSelectedCell();
		if ( selectedCell == null ) {
			return;
		}
		if (selectedCell.isRight()) {
			return;
		}
		
		var sibling = this.getSiblingCell(selectedCell);
		var cells = [];
		var device_id = this.device_id;
		if ( sibling != null ) {
			sibling.position = "LEFT";
			selectedCell.position = "RIGHT";
			cells.push(sibling);
			cells.push(selectedCell);
		} else if ( selectedCell.isLeft() ) {
			selectedCell.position = "BANNER";
			cells.push(selectedCell);
		} else if ( selectedCell.isBanner() ) {
			selectedCell.position = "RIGHT";
			cells.push(selectedCell);
		}
		
		if ( cells.length > 0 ) {
			Server.device_move_cells(device_id, cells);
			App.View.showDragDrop();
		}
		
	};
	
	this.removeSelectedCell = function() {
		Log.d("DeviceLayout.moveSelectedCellRight()");
		var cell = this.getSelectedCell();
		if ( cell == null ) {
			return;
		}
		var cell_index = this.cells.content.indexOf(cell);
		this.cells.content.splice(cell_index, 1);
		var device_id = this.device_id;
		if (cell.type == "contact" ) {
			Server.device_hide_contact(device_id, cell.page_id, cell.contact_id);
			
		} else if ( cell.type == "app" ) {
			Server.device_hide_app(device_id, cell.page_id, cell.app_id);
		}
		var cells = [];
		this.fixAnyEmptyRows(cells);
		if ( cells.length > 0 ) {
			Server.device_move_cells(device_id, cells);
		}
		App.View.showDragDrop();

	};
	
	this.fixAnyEmptyRows = function(cells) {
		
		for ( var page_id = 1; page_id < 10; page_id ++ ) {
			var max_row = 0;
			var min_row = 10;
			for (var index=0; index<this.cells.content.length; index++) {
				var cell = this.cells.content[index];
				if ( cell.page_id != page_id ) {
					continue;
				}
				if (cell.row > max_row ) {
					max_row = cell.row;
				}
				if ( cell.row < min_row ) {
					min_row = cell.row;
				}
			}
			
			var adjust = 0;
			for (var row=1; row<=max_row; row++) {
				var cells_for_row = this.getCellsForRow(row, page_id); 
				if ( cells_for_row.length == 0 ) {
					// everything after this row needs to be decreased by 1
					adjust = new Number(adjust) + 1;
				} else {
					if ( adjust > 0 ) {
						for(var index=0; index<cells_for_row.length; index++) {
							cells_for_row[index].row = new Number(cells_for_row[index].row) - adjust;
							cells.push(cells_for_row[index]);
						}
					}
				}
			}
		}
		
		
	};

    _.extend(this, data);
    this.init();
				
};

var DeviceCell = function(data) {
	
    _.extend(this, data);

    this.getId = function() {
    	if ( this.type == "contact" ) {
    		return this.contact_id;
    	} else if ( this.type == "app") {
    		return this.app_id;
    	} else {
    		return this.id;
    	}
    };
    
    this.render = function(container_element) {

		var sel_class = this.isSelected()  ? "dnd_cell_selected" : "";
		
		var img = "/app/themes/basic/" + this.type + "s/" + this.getId() + ".png";
		
		var showImages = false;
    	if ( this.isBanner() ) {
    		
    		if ( showImages ) {
				var content = $("<td cell_id='" + this.getId() + "' class='dnd_cell dnd_banner_cell " + sel_class + "' colspan='2'><img class='dnd_cell_banner_image' src='" + img + "'/></td>");
				container_element.append(content); // $(".dragdrop_table").append(content);
    		} else {
				var content = $("<td cell_id='" + this.getId() + "' class='dnd_cell dnd_banner_cell " + sel_class + "' colspan='2'>" + this.getId() + "</td>");
				container_element.append(content); // $(".dragdrop_table").append(content);
    			
    		}
			
    	} else if ( this.isLeft() ) {
    		if ( showImages ) {
	    		var content = $("<td cell_id='" + this.getId() + "' class='dnd_cell dnd_left_cell " + sel_class + "' ><img class='dnd_cell_left_image' src='" + img + "'/></td>");
				container_element.append(content);
    		} else {
	    		var content = $("<td cell_id='" + this.getId() + "' class='dnd_cell dnd_left_cell " + sel_class + "' >" + this.getId() + "</td>");
				container_element.append(content);
    		}
			
    	} else if ( this.isRight() ) {
    		if ( showImages ) {
				var content = $("<td cell_id='" + this.getId() + "' class='dnd_cell dnd_right_cell " + sel_class + "'><img class='dnd_cell_right_image' src='" + img + "'/></td>");
				container_element.append(content);
    		} else {
				var content = $("<td cell_id='" + this.getId() + "' class='dnd_cell dnd_right_cell " + sel_class + "'>" + this.getId() + "</td>");
				container_element.append(content);
    			
    		}
			
    	}
    };
    
    this.click = function() {
    	if ( this.is_selected = true ) {
    		this.is_selected = false;
    	} else {
    		
    	}
    };
    
    this.isSelected = function() {
    	return this.is_selected;
    };
    
    this.setSelected = function(value) {
    	this.is_selected = value;
    };
    
	this.isBanner = function() {
		return this.position.toLowerCase() == "banner";
	};

	this.isRight = function() {
		return this.position.toLowerCase() == "right";
	};

	this.isLeft = function() {
		return this.position.toLowerCase() == "left";
	};

    
    
};