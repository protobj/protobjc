package ts

const EnumSchemaTemplate = `
export class ${class}Schema implements Schema{


    writeTo(output: Output, message: any, polymorphic: boolean): void {
       ${class}Schema.writeTo(output,message as ${class}, polymorphic); 	
    }

    static writeTo(output: Output, message:${class},polymorphic:boolean) {
${writeBody}
    }

    static writeToWithFieldNumber(fieldNum:number,output: Output, message:${class},polymorphic:boolean) {
${writeWithFieldNumberBody}
    }

    mergeFrom(input: Input, message: any):any {
        return ${class}Schema.mergeFrom(input,message as ${class})
    }
    
    static mergeFrom(input: Input, message: ${class}):${class} {
${readBody}
    }
}

`
const MessageSchemaTemplate = `
export class ${class}Schema implements Schema {
               
	writeTo(output:Output, message:any, polymorphic:boolean){
		${class}Schema.writeTo(output, message as ${class}, polymorphic);
	}
	
	static writeTo(output:Output, message:${class}, polymorphic:boolean)  {
		if (polymorphic) {
			output.writeI32_Packed(${messageIndex});
		}
${writeBody}
	}

	mergeFrom(input:Input, message:any):any {
		return ${class}Schema.mergeFrom(input, message as ${class});
	}
	
	static mergeFrom(input:Input, msg:${class}):${class} {
		let fieldNumber = input.readFieldNumber();
		if (fieldNumber == 0) {
			return null;
		}
		if (msg == null) {
			msg = new ${class}();
		}
		const message:${class} = msg;
		do {
			switch (fieldNumber) {
${readBody}
				default: {
					input.handleUnknownField(fieldNumber);
					break;
				}
			}
				
		} while ((fieldNumber = input.readFieldNumber()) != 0);
		return message;
	}
}
`
