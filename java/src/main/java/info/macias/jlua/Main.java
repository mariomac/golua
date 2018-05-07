package info.macias.jlua;

import org.luaj.vm2.Globals;
import org.luaj.vm2.LuaString;
import org.luaj.vm2.LuaValue;
import org.luaj.vm2.Varargs;
import org.luaj.vm2.lib.OneArgFunction;
import org.luaj.vm2.lib.jse.JsePlatform;

public class Main {

	private static class HostPrintFunc extends OneArgFunction {
		@Override
		public LuaValue call(LuaValue arg) {
			System.out.println(arg.checkjstring());
			return NIL;
		}
	}

	public static void main(String[] main) {
		Globals globals = JsePlatform.standardGlobals();
		// Load external script on Lua Globals context
		globals.loadfile("../lua/hello.lua").call();
		// Registering java-hosted function into the global context
		globals.set("host_print", new HostPrintFunc());
		// Invoking global symbol "sayhello" from the script
		LuaValue length = globals.get("sayhello").call(LuaString.valueOf("my friend"));
		// Catching returned values
		System.out.println("the length of the argument is " + length.checkint());
	}
}
