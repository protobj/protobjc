package java

const EnumSchemaTemplate = `
public class ${class}Schema implements Schema {
                
	@Override
	public void writeTo(Output output, Object message, boolean polymorphic) throws IOException {
		writeTo(output, (${class}) message, polymorphic);
	}

	public static void writeTo(Output output, ${class} message, boolean polymorphic) throws IOException {
${writeBody}
	}
	public static void writeWithFieldNumber(int fieldNum, Output output, ${class} message) throws IOException {
${writeWithFieldNumberBody}
	}

	@Override
	public Object mergeFrom(Input input, Object message) throws IOException {
		return mergeFrom(input, (${class}) message);
	}

	public static ${class} mergeFrom(Input input, ${class} message) throws IOException {
${readBody}
	}

}
`
const MessageSchemaTemplate = `
public class ${class}Schema implements Schema {
               
	@Override
	public void writeTo(Output output, Object message, boolean polymorphic) throws IOException {
		writeTo(output, (${class}) message, polymorphic);
	}
	
	public static void writeTo(Output output, ${class} message, boolean polymorphic) throws IOException {
		if (polymorphic) {
			output.writeI32(${messageIndex});
		}
${writeBody}
	}
	
	@Override
	public Object mergeFrom(Input input, Object message) throws IOException {
		return mergeFrom(input, (${class}) message);
	}
	
	public static ${class} mergeFrom(Input input, ${class} msg) throws IOException {
		int fieldNumber = input.readFieldNumber();
		if (fieldNumber == 0) {
			return null;
		}
		if (msg == null) {
			msg = new ${class}();
		}
		${class} message = msg;
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
